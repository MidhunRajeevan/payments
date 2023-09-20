package payment

import (
	"MidhunRajeevan/payments/app/gateways"
	"MidhunRajeevan/payments/app/util"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	pgx "github.com/jackc/pgx/v4"
)

func GenerateTransactionToken(w http.ResponseWriter, r *http.Request) {

	var record TransactionRequest
	// Parse the JSON request
	var paymentReq PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&paymentReq)
	if err != nil {
		util.BadRequest(&w, "Failed to parse JSON request")
		return
	}

	sourceSystemID := chi.URLParam(r, "source-systems-did")
	if !validateRequest(w, paymentReq, sourceSystemID) {
		return
	}

	uuid, token := util.GenerateTransactionToken()

	paymentReq.Ewallet.PIN, err = util.Encrypt(paymentReq.Ewallet.PIN)
	if err != nil {
		log.Println("Encrypt Error for paymentReq.Ewallet.PIN", err.Error())
	}

	jsonString, err := json.Marshal(paymentReq)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	record.Payload = string(jsonString)
	record.TransactionReference = uuid
	record.Token = token
	record.DID = sourceSystemID

	record, err = insertRequest(record)
	if err != nil {
		log.Println("Insert request Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	lstGateways, err := gateways.SelectGatewaysByDid(sourceSystemID)
	if err == pgx.ErrNoRows {
		util.NotFound(&w, "not_found")
		return
	} else if err != nil {
		log.Println("Select Source System Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	res := TransactionRequestResponseMapping(lstGateways, record)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func CreateInvoice(w http.ResponseWriter, r *http.Request) {

	var req InvoiceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.BadRequest(&w, "Failed to parse JSON request")
		return
	}
	sourceSystemID := chi.URLParam(r, "source-systems-did")
	if !validateRequest(w, req, sourceSystemID) {
		return
	}

	//Get company details(register to RMS if not already registered)
	company, ok := getCompanies(w, req)
	if !ok {
		log.Println("get customer error:", err.Error())
		return
	}

	//Upsert companies contact
	contacts, ok := companyContactMapping(w, company, req)
	if !ok {
		log.Println("company contact mapping error:", err.Error())
		return
	}
	_, err = upsertContactTable(contacts)
	if err != nil {
		log.Println("Insert request Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	//create RMS sales order
	invoiceNumber, ok := generateSalesOrderRMS(req, company)
	if !ok {
		log.Println("Sales Order Creation Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}
	invoiceEntity, ok := InvoiceEntityMapping(w, req, sourceSystemID, company.ID, invoiceNumber)
	if !ok {
		log.Println("Invoice Entity Mapping error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	log.Println(invoiceEntity)
	invoiceTable, err := insertInvoiceTable(invoiceEntity)
	if err != nil {
		log.Println("Insert request Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoiceTable)

}

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

func ProcessPayment(w http.ResponseWriter, r *http.Request) {

	var record TransactionRequest
	// Parse the JSON request
	var paymentReq PaymentRequest
	err := json.NewDecoder(r.Body).Decode(&paymentReq)
	if err != nil {
		util.BadRequest(&w, "Failed to parse JSON request")
		return
	}

	sourceSystemID := chi.URLParam(r, "source-systems-ID")
	if !validatePaymentRequest(w, paymentReq, sourceSystemID) {
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

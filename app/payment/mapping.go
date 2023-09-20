package payment

import (
	"MidhunRajeevan/payments/app/gateways"
	"MidhunRajeevan/payments/app/util"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func TransactionRequestResponseMapping(lstGateways []gateways.PaymentGateway, record TransactionRequest) TransactionRequestResponse {
	var response TransactionRequestResponse

	response.ID = record.ID
	response.DID = record.DID
	response.Token = record.Token
	response.TransactionReference = record.TransactionReference
	response.Status = record.Status
	response.Attempts = record.Attempts
	response.Message = record.Message
	response.IsExpired = record.IsExpired
	response.PaymentInitiatedAt = record.PaymentInitiatedAt
	response.CreatedAt = record.CreatedAt
	response.UpdatedAt = record.UpdatedAt

	for _, channel := range lstGateways {
		var gatewayHeader GatewayHeader
		gatewayHeader.ID = channel.ID
		gatewayHeader.StandardName = channel.StandardName
		gatewayHeader.NameEn = channel.NameEn
		gatewayHeader.NameAr = channel.NameAr
		response.Channels = append(response.Channels, gatewayHeader)
	}

	return response
}

func companyContactMapping(w http.ResponseWriter, company Companies, req InvoiceRequest) (ContactTable, bool) {
	var contacts ContactTable
	contacts.ContactableID = company.ID
	contacts.ContactableType = "company"
	contacts.ContactableKey = company.TradeLicenseNumber
	if req.Company.Contact.Addresses == nil {
		contacts.Addresses = "[]"
	} else {
		addresses, err := json.Marshal(req.Company.Contact.Addresses)
		if err != nil {
			log.Println("json marshel error for Company.Contact.Address:", err)
			util.InternalServerError(&w, "contact_support")
			return contacts, false
		}
		contacts.Addresses = string(addresses)
	}
	if req.Company.Contact.Emails == nil {
		contacts.Emails = "[]"
	} else {
		emails, err := json.Marshal(req.Company.Contact.Emails)
		if err != nil {
			log.Println("json marshel error for Company.Contact.Emails:", err)
			util.InternalServerError(&w, "contact_support")
			return contacts, false
		}
		contacts.Emails = string(emails)
	}

	if req.Company.Contact.Phones == nil {
		contacts.Emails = "[]"
	} else {
		Phones, err := json.Marshal(req.Company.Contact.Phones)
		if err != nil {
			log.Println("json marshel error for Company.Contact.Phones:", err)
			util.InternalServerError(&w, "contact_support")
			return contacts, false
		}
		contacts.Phones = string(Phones)
	}
	contacts.Attachments = "[]"

	return contacts, true
}

func InvoiceEntityMapping(w http.ResponseWriter, req InvoiceRequest, did string, CompanyID int64,
	invoiceNumber string) (InvoiceTable, bool) {

	var invoice InvoiceTable

	invoice.Did = did
	invoice.CompanyID = CompanyID
	invoice.Number = invoiceNumber
	invoice.IssuedDate = req.Invoice.IssuedDate
	invoice.LastDate = req.Invoice.LastDate
	invoice.Description = req.Invoice.Description
	invoice.AmountValue = req.Invoice.Amount.Value
	invoice.AmountCurrency = req.Invoice.Amount.Currency
	invoice.InvoiceReferenceNumber = uuid.New().String()
	invoice.PurchaseOrderNo = invoice.InvoiceReferenceNumber
	requester, err := json.Marshal(req.Person)
	if err != nil {
		log.Println("json marshel error for requester:", err)
		util.InternalServerError(&w, "contact_support")
		return invoice, false
	}
	invoice.Requester = string(requester)

	return invoice, true

}

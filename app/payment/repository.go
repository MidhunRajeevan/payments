package payment

import (
	"MidhunRajeevan/payments/config"
	"context"
	"log"
)

func execAndScanRequestTable(statement string, params []interface{}) (TransactionRequest, error) {
	var w TransactionRequest
	err := config.DB.QueryRow(context.Background(), statement, params...).Scan(&w.ID,
		&w.DID, &w.Payload, &w.Token, &w.TransactionReference, &w.Status, &w.Attempts,
		&w.Message, &w.IsExpired, &w.CreatedAt, &w.UpdatedAt)
	if err != nil {
		log.Println("Error in Database Query.")
		log.Println("Statement :", statement)
		log.Println("Error :", err.Error())
	}
	return w, err
}

func execAndScanCompaniesTable(statement string, params []interface{}) (Companies, error) {
	var w Companies
	err := config.DB.QueryRow(context.Background(), statement, params...).Scan(&w.ID,
		&w.Name.En, &w.Name.Ar, &w.CustGroup, &w.Email, &w.TradeLicenseNumber, &w.TrafficFileNumber, &w.Account, &w.CreatedAt, &w.UpdatedAt)
	if err != nil {
		log.Println("Error in Database Query.")
		log.Println("Statement :", statement)
		log.Println("Error :", err.Error())
	}
	return w, err
}

func execAndScanContactTable(statement string, params []interface{}) (ContactTable, error) {
	var w ContactTable
	err := config.DB.QueryRow(context.Background(), statement, params...).Scan(&w.ID,
		&w.ContactableType, &w.ContactableID, &w.ContactableKey, &w.Addresses, &w.Phones, &w.Emails, &w.Attachments, &w.CreatedAt, &w.UpdatedAt)
	if err != nil {
		log.Println("Error in Database Query.")
		log.Println("Statement :", statement)
		log.Println("Error :", err.Error())
	}
	return w, err
}

func execAndScanInvoiceTable(statement string, params []interface{}) (InvoiceTable, error) {
	var w InvoiceTable
	err := config.DB.QueryRow(context.Background(), statement, params...).Scan(&w.ID,
		&w.Did, &w.CompanyID, &w.Number, &w.IssuedDate, &w.LastDate, &w.Description, &w.AmountValue,
		&w.AmountCurrency, &w.PaymentStatus, &w.PaymentDate, &w.InvoiceReferenceNumber, &w.PaymentReferenceNumber,
		&w.PurchaseOrderNo, &w.DocumentDate, &w.RMSInvoiceDate, &w.Requester, &w.CreatedAt, &w.UpdatedAt)
	if err != nil {
		log.Println("Error in Database Query.")
		log.Println("Statement :", statement)
		log.Println("Error :", err.Error())
	}
	return w, err
}
func insertRequest(r TransactionRequest) (TransactionRequest, error) {
	statement := `
	insert into requests
	(did, payload, token, transactionreference)
	values ($1, $2, $3, $4)
	returning
		id, did, payload, token, transactionreference, status,
		attempts, coalesce(message, '') as message, is_expired,
	  	created_at, updated_at`
	params := make([]interface{}, 0)
	params = append(params,
		r.DID, r.Payload, r.Token, r.TransactionReference)
	return execAndScanRequestTable(statement, params)
}

// Companies

func upsertCompaniesTable(r Companies) (Companies, error) {
	statement := `
	insert into companies
	(name_en, name_ar, custgroup, email, trade_license_number,traffic_file_number, account)
	values ($1, $2, $3, $4, $5,$6,$7)
	on conflict (trade_license_number) do update
		set name_en = $1,
		name_ar = $2,
		custgroup = $3,
		email = $4,
		account = $7,
		traffic_file_number=$6,
		updated_at = current_timestamp
		where companies.trade_license_number = $5
	returning
		id,name_en, name_ar, custgroup, email, trade_license_number,traffic_file_number, account,
	  	created_at, updated_at`
	params := make([]interface{}, 0)
	params = append(params,
		r.Name.En, r.Name.Ar, r.CustGroup, r.Email, r.TradeLicenseNumber, r.TrafficFileNumber, r.Account)
	return execAndScanCompaniesTable(statement, params)
}

func selectCompanies(tradeLicenseNumber string) (Companies, error) {
	statement := `
	select
	id,name_en, name_ar, custgroup, email, trade_license_number,traffic_file_number, account,
	created_at, updated_at
	from companies
	where trade_license_number = $1`
	params := make([]interface{}, 0)
	params = append(params, tradeLicenseNumber)
	return execAndScanCompaniesTable(statement, params)
}

// contacts

func upsertContactTable(r ContactTable) (ContactTable, error) {
	statement := `
	insert into contacts
	(contactable_type, contactable_id, contactable_key, addresses, phone_numbers,emails, attachments)
	values ($1, $2, $3, $4, $5,$6,$7)
	on conflict (contactable_type,contactable_id) do update
		set contactable_key = $3,
		addresses = $4,
		phone_numbers = $5,
		emails = $6,
		attachments = $7,
		updated_at = current_timestamp
		where contacts.contactable_type = $1 and contacts.contactable_id =$2
	returning
		id,contactable_type, contactable_id, contactable_key, addresses, phone_numbers,emails, attachments,
	  	created_at, updated_at`
	params := make([]interface{}, 0)
	params = append(params,
		r.ContactableType, r.ContactableID, r.ContactableKey, r.Addresses, r.Phones, r.Emails, r.Attachments)
	return execAndScanContactTable(statement, params)
}

func insertInvoiceTable(r InvoiceTable) (InvoiceTable, error) {
	statement := `
	insert into invoices
	(did, company_id, "number", issued_date, last_date, description, 
	amount_value, amount_currency, invoice_reference_number ,purchase_order_form_no,
	document_date, rms_invoice_date, requester)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	returning
		id,did, company_id, "number", issued_date, last_date, description, amount_value, amount_currency, payment_status, payment_date,invoice_reference_number,payment_reference_number, purchase_order_form_no,
		document_date, rms_invoice_date, requester, created_at, updated_at`
	params := make([]interface{}, 0)
	params = append(params,
		r.Did, r.CompanyID, r.Number, r.IssuedDate, r.LastDate, r.Description,
		r.AmountValue, r.AmountCurrency, r.InvoiceReferenceNumber, r.PurchaseOrderNo,
		r.DocumentDate, r.RMSInvoiceDate, r.Requester)
	return execAndScanInvoiceTable(statement, params)
}

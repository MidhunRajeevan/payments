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

package payment

import (
	"MidhunRajeevan/payments/app/sources"
	"MidhunRajeevan/payments/app/util"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	pgx "github.com/jackc/pgx/v4"
)

func validatePaymentRequest(w http.ResponseWriter, req PaymentRequest, sourceSystemID string) bool {

	validate := validator.New()
	err := validate.Struct(req)

	if err != nil {
		util.BadRequest(&w, "invalid_request")
		return false
	}

	// check if the source system ID passed in the URL and in request body are same
	if sourceSystemID != req.Source.DID {
		util.BadRequest(&w, "The source system ID does not match")
		return false
	}

	// check if the source system exist in the DB
	_, err = sources.SelectSource(sourceSystemID)
	if err == pgx.ErrNoRows {
		util.NotFound(&w, "not_found")
		return false
	} else if err != nil {
		log.Println("Select Source System Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return false
	}

	return true

}

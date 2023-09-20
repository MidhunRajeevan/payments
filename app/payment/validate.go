package payment

import (
	"MidhunRajeevan/payments/app/sources"
	"MidhunRajeevan/payments/app/util"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	pgx "github.com/jackc/pgx/v4"
)

func validateRequest(w http.ResponseWriter, req interface{}, sourceSystemID string) bool {

	validate := validator.New()
	err := validate.Struct(req)

	if err != nil {
		util.BadRequest(&w, "invalid_request")
		return false
	}

	// check if the source system exist in the DB
	source, err := sources.SelectSource(sourceSystemID)
	if err == pgx.ErrNoRows {
		util.NotFound(&w, "The Source system "+sourceSystemID+" not found")
		return false
	} else if err != nil {
		log.Println("Select Source System Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return false
	}
	if !source.IsActive {
		util.NotFound(&w, "Source "+sourceSystemID+"is not active")
		return false
	}

	return true

}

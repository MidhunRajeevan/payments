package sources

import (
	"MidhunRajeevan/payments/app/gateways"
	"MidhunRajeevan/payments/app/util"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	pgx "github.com/jackc/pgx/v4"
)

func validateSourceSystem(w http.ResponseWriter, sourceSystemID string) (SourceSystem, bool) {

	// check if the source system exist in the DB
	source, err := SelectSource(sourceSystemID)
	if err == pgx.ErrNoRows {
		util.NotFound(&w, "Source Did"+sourceSystemID+"does not exist")
		return source, false
	} else if err != nil {
		log.Println("Select Source System Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return source, false
	}
	if !source.IsActive {
		util.NotFound(&w, "Source "+sourceSystemID+"is not active")
		return source, false
	}

	return source, true
}

func validateSourceAndGateway(w http.ResponseWriter, req SourceSystemGateway, sourceSystemID string) bool {

	validate := validator.New()
	err := validate.Struct(req)

	if err != nil {
		util.BadRequest(&w, "invalid_request")
		return false
	}

	// check if the source system ID passed in the URL and in request body are same
	if sourceSystemID != req.DID {
		util.BadRequest(&w, "The source system ID does not match")
		return false
	}

	// check if the source system exist in the DB
	_, ok := validateSourceSystem(w, sourceSystemID)
	if !ok {
		return false
	}

	gateway, err := gateways.SelectGatewayByCode(req.GatewayCode)
	if err == pgx.ErrNoRows {
		util.NotFound(&w, "Gateway "+req.GatewayCode+"does not exist")
		return false
	} else if err != nil {
		log.Println("Select gateway Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return false
	}
	if !gateway.IsActive {
		util.NotFound(&w, "Gateway "+req.GatewayCode+"is not active")
		return false
	}

	return true

}

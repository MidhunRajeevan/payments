package sources

import (
	"MidhunRajeevan/payments/app/util"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func mapGatewayToSource(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request

	var req SourceSystemGateway
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.BadRequest(&w, "Failed to parse JSON request")
		return
	}

	sourceSystemID := chi.URLParam(r, "source-systems-ID")
	if !validatePaymentRequest(w, req, sourceSystemID) {
		return
	}
	record, err := insertSourceSystemGateway(req)
	if err != nil {
		log.Println("Insert request Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)

}

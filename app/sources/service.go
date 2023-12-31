package sources

import (
	"MidhunRajeevan/payments/app/gateways"
	"MidhunRajeevan/payments/app/util"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
	pgx "github.com/jackc/pgx/v4"
)

func createSource(w http.ResponseWriter, r *http.Request) {

	// Parse the JSON request
	var req SourceSystem
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.BadRequest(&w, "Failed to parse JSON request")
		return
	}

	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		util.BadRequest(&w, "invalid_request")
		return
	}

	record, err := upsertSource(req)
	if err != nil {
		log.Println("Insert request Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func getSources(w http.ResponseWriter, r *http.Request) {

	record, err := SelectSources()
	if err != nil {
		log.Println("Select Sources Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func getSourceByID(w http.ResponseWriter, r *http.Request) {

	sourceSystemID := chi.URLParam(r, "source-systems-did")
	record, err := selectSourceByCode(sourceSystemID)
	if err != nil {
		log.Println("Select Sources Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func updateSourceByID(w http.ResponseWriter, r *http.Request) {

	// Parse the JSON request
	var req SourceSystem
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.BadRequest(&w, "Failed to parse JSON request")
		return
	}
	sourceSystemID := chi.URLParam(r, "source-systems-did")
	req.Did = sourceSystemID
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		util.BadRequest(&w, "invalid_request")
		return
	}

	record, err := upsertSource(req)
	if err != nil {
		log.Println("Insert request Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func activateSourceByID(w http.ResponseWriter, r *http.Request) {

	sourceSystemID := chi.URLParam(r, "source-systems-did")
	record, err := selectSourceByCode(sourceSystemID)
	if err != nil {
		log.Println("Select Sources Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}
	record.IsActive = true

	res, err := updateSourceStatus(record)
	if err != nil {
		log.Println("Update Source status Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func deactivateSourceByID(w http.ResponseWriter, r *http.Request) {

	sourceSystemID := chi.URLParam(r, "source-systems-did")
	record, err := selectSourceByCode(sourceSystemID)
	if err != nil {
		log.Println("Select Sources Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}
	record.IsActive = false

	res, err := updateSourceStatus(record)
	if err != nil {
		log.Println("Update Source status Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func mapGatewayToSource(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request

	var req SourceSystemGateway
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.BadRequest(&w, "Failed to parse JSON request")
		return
	}

	sourceSystemID := chi.URLParam(r, "source-systems-did")
	if !validateSourceAndGateway(w, req, sourceSystemID) {
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

func getSourceGatewayMapping(w http.ResponseWriter, r *http.Request) {
	// Parse the JSON request

	sourceSystemID := chi.URLParam(r, "source-systems-did")

	lstGateways, err := gateways.SelectGatewaysByDid(sourceSystemID)
	if err == pgx.ErrNoRows {
		util.NotFound(&w, "not_found")
		return
	} else if err != nil {
		log.Println("Select Source System Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lstGateways)

}

func activateSourceGatewayMapping(w http.ResponseWriter, r *http.Request) {

	var req SourceSystemGateway
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.BadRequest(&w, "Failed to parse JSON request")
		return
	}

	sourceSystemID := chi.URLParam(r, "source-systems-did")
	if !validateSourceAndGateway(w, req, sourceSystemID) {
		return
	}
	req.IsActive = true

	res, err := updateSourceSystemGatewayStatus(req)
	if err != nil {
		log.Println("Update Source-System Gateway status Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func deActivateSourceGatewayMapping(w http.ResponseWriter, r *http.Request) {

	var req SourceSystemGateway
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.BadRequest(&w, "Failed to parse JSON request")
		return
	}

	sourceSystemID := chi.URLParam(r, "source-systems-did")
	if !validateSourceAndGateway(w, req, sourceSystemID) {
		return
	}
	req.IsActive = false

	res, err := updateSourceSystemGatewayStatus(req)
	if err != nil {
		log.Println("Update Source-System Gateway status Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

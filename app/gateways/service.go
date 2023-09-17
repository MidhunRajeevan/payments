package gateways

import (
	"MidhunRajeevan/payments/app/util"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

func createGateway(w http.ResponseWriter, r *http.Request) {

	// Parse the JSON request
	var req PaymentGateway
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

	record, err := upsertGateway(req)
	if err != nil {
		log.Println("Insert request Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func getGateways(w http.ResponseWriter, r *http.Request) {

	record, err := SelectGateways()
	if err != nil {
		log.Println("Select Gateways Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func getGatewayByID(w http.ResponseWriter, r *http.Request) {

	gatewayCode := chi.URLParam(r, "gateway-code")
	record, err := SelectGatewayByCode(gatewayCode)
	if err != nil {
		log.Println("Select Gateways Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func updateGatewayByID(w http.ResponseWriter, r *http.Request) {

	// Parse the JSON request
	var req PaymentGateway
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.BadRequest(&w, "Failed to parse JSON request")
		return
	}
	gatewayCode := chi.URLParam(r, "gateway-code")
	req.Code = gatewayCode
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		util.BadRequest(&w, "invalid_request")
		return
	}

	record, err := upsertGateway(req)
	if err != nil {
		log.Println("Insert request Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func activateGatewayByID(w http.ResponseWriter, r *http.Request) {

	gatewayCode := chi.URLParam(r, "gateway-code")
	record, err := SelectGatewayByCode(gatewayCode)
	if err != nil {
		log.Println("Select Gateways Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}
	record.IsActive = true

	res, err := updateGatewayStatus(record)
	if err != nil {
		log.Println("Update gateway status Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func deactivateGatewayByID(w http.ResponseWriter, r *http.Request) {

	gatewayCode := chi.URLParam(r, "gateway-code")
	record, err := SelectGatewayByCode(gatewayCode)
	if err != nil {
		log.Println("Select Gateways Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}
	record.IsActive = false

	res, err := updateGatewayStatus(record)
	if err != nil {
		log.Println("Update gateway status Error:", err.Error())
		util.InternalServerError(&w, "contact_support")
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

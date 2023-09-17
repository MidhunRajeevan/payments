package gateways

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Post("/", createGateway)
	mux.Get("/", getGateways)
	mux.Get("/{gateway-code}", getGatewayByID)
	mux.Put("/{gateway-code}", updateGatewayByID)
	mux.Post("/{gateway-code}/events/activate", activateGatewayByID)
	mux.Post("/{gateway-code}/events/deactivate", deactivateGatewayByID)

	return mux
}

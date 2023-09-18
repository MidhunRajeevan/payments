package sources

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {
	mux := chi.NewRouter()
	mux.Post("/", createSource)
	mux.Get("/", getSources)
	mux.Get("/{source-systems-did}", getSourceByID)
	mux.Put("/{source-systems-did}", updateSourceByID)
	mux.Post("/{source-systems-did}/events/activate", activateSourceByID)
	mux.Post("/{source-systems-did}/events/deactivate", deactivateSourceByID)
	mux.Post("/{source-systems-did}/gateway-mapping", mapGatewayToSource)
	mux.Post("/{source-systems-did}/gateway-mapping/events/activate", activateSourceGatewayMapping)
	mux.Post("/{source-systems-did}/gateway-mapping/events/deActivate", deActivateSourceGatewayMapping)

	return mux
}

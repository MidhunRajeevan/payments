package sources

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {
	mux := chi.NewRouter()
	mux.Post("/{source-systems-did}/gateway-mapping", mapGatewayToSource)

	// Define routes specific to source package

	return mux
}

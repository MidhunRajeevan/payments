package payment

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {
	mux := chi.NewRouter()
	mux.Post("/source-systems/{source-systems-ID}", ProcessPayment)

	return mux
}

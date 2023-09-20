package payment

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {
	mux := chi.NewRouter()
	mux.Post("/token", GenerateTransactionToken)

	return mux
}

func InvoiceRoutes() http.Handler {
	mux := chi.NewRouter()
	mux.Post("/", CreateInvoice)

	return mux
}

package routes

import (
	"MidhunRajeevan/payments/app/gateways"
	"MidhunRajeevan/payments/app/payment"
	"MidhunRajeevan/payments/app/sources"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {
	mux := chi.NewRouter()
	mux.Mount("/gateways", gateways.Routes())
	mux.Mount("/source-systems/{source-systems-did}/payments", payment.Routes())
	mux.Mount("/source-systems/{source-systems-did}/invoices", payment.InvoiceRoutes())
	mux.Mount("/source-systems", sources.Routes())
	mux.Get("/", Index)
	return mux
}

// Index API
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{ "app": "rta-payments" }`))
}

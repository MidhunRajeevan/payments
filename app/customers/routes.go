package customers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", Index)
	mux.Post("/", Index)
	// Subrouters:
	mux.Route("/{customerID}", func(mux chi.Router) {
		mux.Get("/", Index)
		mux.Put("/", Index)
		mux.Delete("/", Index)
	})

	return mux
}

// Index API
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{ "app": "rta-payments" }`))
}

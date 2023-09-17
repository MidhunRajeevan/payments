package util

import (
	"encoding/json"
	"net/http"
)

// BadRequest response
func BadRequest(w *http.ResponseWriter, msg string) {
	error := Error{Code: 400, Message: msg}
	(*w).WriteHeader(http.StatusBadRequest)
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(error)
}

// Conflict response
func Conflict(w *http.ResponseWriter, msg string) {
	error := Error{Code: 409, Message: msg}
	(*w).WriteHeader(http.StatusConflict)
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(error)
}

// NotFound response
func NotFound(w *http.ResponseWriter, msg string) {
	error := Error{Code: 404, Message: msg}
	(*w).WriteHeader(http.StatusNotFound)
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(error)
}

// MethodNotAllowed response
func MethodNotAllowed(w *http.ResponseWriter, msg string) {
	error := Error{Code: 405, Message: msg}
	(*w).WriteHeader(http.StatusMethodNotAllowed)
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(error)
}

// InternalServerError response
func InternalServerError(w *http.ResponseWriter, msg string) {
	error := Error{Code: 500, Message: msg}
	(*w).WriteHeader(http.StatusInternalServerError)
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(error)
}

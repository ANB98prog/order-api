package response

import (
	"encoding/json"
	"net/http"
)

func jsonResponse[T any](w http.ResponseWriter, data T, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func jsonErrorResponse[T any](w http.ResponseWriter, data T, status int) {
	w.Header().Set("Content-Type", "application/json")
	message, _ := json.Marshal(data)
	http.Error(w, string(message), status)
}

func OK[T any](w http.ResponseWriter, data T)       { jsonResponse(w, data, http.StatusOK) }
func Created[T any](w http.ResponseWriter, data T)  { jsonResponse(w, data, http.StatusCreated) }
func NotFound[T any](w http.ResponseWriter, data T) { jsonErrorResponse(w, data, http.StatusNotFound) }
func BadRequest[T any](w http.ResponseWriter, data T) {
	jsonErrorResponse(w, data, http.StatusBadRequest)
}
func Unauthorized[T any](w http.ResponseWriter, data T) {
	jsonErrorResponse(w, data, http.StatusUnauthorized)
}
func InternalServerError[T any](w http.ResponseWriter, data T) {
	jsonErrorResponse(w, data, http.StatusInternalServerError)
}

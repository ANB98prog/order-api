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

func OK[T any](w http.ResponseWriter, data T)       { jsonResponse(w, data, http.StatusOK) }
func Created[T any](w http.ResponseWriter, data T)  { jsonResponse(w, data, http.StatusCreated) }
func NotFound[T any](w http.ResponseWriter, data T) { jsonResponse(w, data, http.StatusNotFound) }
func BadRequest(w http.ResponseWriter, err error)   { jsonResponse(w, err, http.StatusBadRequest) }
func InternalServerError[T any](w http.ResponseWriter, data T) {
	jsonResponse(w, data, http.StatusInternalServerError)
}

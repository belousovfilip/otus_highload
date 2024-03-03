package response

import (
	"encoding/json"
	"net/http"
	"otus_highload/internal/lib/appErrors"
)

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Errors struct {
	Code   int               `json:"code"`
	Errors map[string]string `json:"errors"`
}

func Json(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
}

func JsonMessage(w http.ResponseWriter, code int, msg string) {
	Json(w, code)
	json.NewEncoder(w).Encode(Message{Message: msg, Code: code})
}

func JsonSuccess(w http.ResponseWriter, v any) {
	Json(w, http.StatusOK)
	json.NewEncoder(w).Encode(v)
}

func JsonInternalServerError(w http.ResponseWriter) {
	Json(w, http.StatusInternalServerError)
	json.NewEncoder(w).Encode(Message{
		Message: http.StatusText(http.StatusInternalServerError),
		Code:    http.StatusInternalServerError,
	})
}

func JsonValidation(w http.ResponseWriter, errs appErrors.ValidationErrors) {
	Json(w, http.StatusBadRequest)
	json.NewEncoder(w).Encode(Errors{Errors: errs, Code: http.StatusBadRequest})
}

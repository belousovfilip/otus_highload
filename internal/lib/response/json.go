package response

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Message struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Errors struct {
	Code   int               `json:"code"`
	Errors map[string]string `json:"errors"`
}

type Error struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type JSON struct {
	log *slog.Logger
}

func NewJSON(l *slog.Logger) *JSON {
	l = l.With("response", "json")
	return &JSON{log: l}
}

func (r JSON) Encode(w http.ResponseWriter, body any) error {
	return json.NewEncoder(w).Encode(body)
}

func (r JSON) JSONHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
}

func (r JSON) JSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		r.log.With("method", "json-encode").Error(err.Error())
	}
}

func (r JSON) Message(w http.ResponseWriter, code int, msg string) {
	r.JSON(w, code, Message{Message: msg, Code: code})
}

func (r JSON) DefaultHTTPMessage(w http.ResponseWriter, code int) {
	r.JSON(w, code, Message{Message: http.StatusText(code), Code: code})
}

func (r JSON) Success(w http.ResponseWriter, v any) {
	r.JSON(w, http.StatusOK, v)
}

func (r JSON) Error(w http.ResponseWriter, code int, msg string) {
	r.JSON(w, code, Error{Error: msg, Code: code})
}

func (r JSON) ErrorW(w http.ResponseWriter, code int, err error) {
	r.JSON(w, code, Error{Error: err.Error(), Code: code})
}

func (r JSON) SuccessOK(w http.ResponseWriter) {
	r.JSON(w, http.StatusOK, Message{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	})
}

func (r JSON) BadRequestMsg(w http.ResponseWriter, msg string) {
	r.JSON(w, http.StatusBadRequest, Message{
		Code:    http.StatusBadRequest,
		Message: msg,
	})
}

func (r JSON) InternalServerError(w http.ResponseWriter) {
	r.JSON(w, http.StatusInternalServerError, Message{
		Message: http.StatusText(http.StatusInternalServerError),
		Code:    http.StatusInternalServerError,
	})
}

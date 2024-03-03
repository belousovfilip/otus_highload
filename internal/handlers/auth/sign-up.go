package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"otus_highload/internal/lib/appErrors"
	"otus_highload/internal/lib/http/response"
	"otus_highload/internal/services/auth"
)

func (h Handler) SignUp(res http.ResponseWriter, req *http.Request) {
	dto := auth.SignUpDto{}
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil && err != io.EOF {
		h.log.Error(err.Error())
		response.JsonInternalServerError(res)
		return
	}
	err = h.service.SignUp(req.Context(), dto)
	if err, ok := err.(appErrors.ValidationErrors); ok {
		response.JsonValidation(res, err)
		return
	}
	if err != nil {
		h.log.Error(err.Error())
		response.JsonInternalServerError(res)
		return
	}
	response.JsonMessage(res, http.StatusCreated, http.StatusText(http.StatusCreated))
}

package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"otus_highload/internal/lib/appErrors"
	"otus_highload/internal/lib/http/response"
	"otus_highload/internal/services/auth"
)

func (h Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	dto := auth.SignInDto{}
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil && err != io.EOF {
		h.log.Error(err.Error())
		response.JsonInternalServerError(w)
		return
	}
	token, isAuth, err := h.service.SignIn(r.Context(), dto)
	if err, ok := err.(appErrors.ValidationErrors); ok {
		response.JsonValidation(w, err)
		return
	}
	if err != nil {
		h.log.Error(err.Error())
		response.JsonInternalServerError(w)
		return
	}
	if !isAuth {
		response.JsonMessage(w, http.StatusBadRequest, "Credentials do not match")
		return
	}
	response.JsonSuccess(w, struct {
		Token string `json:"token"`
	}{Token: token})
}

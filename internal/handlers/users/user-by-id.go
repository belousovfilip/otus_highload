package users

import (
	"net/http"
	"otus_highload/internal/lib/http/response"
	"strconv"
)

func (h Handler) UserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		response.JsonMessage(w, http.StatusBadRequest, "invalid user id")
		return
	}
	user, err := h.service.UserById(r.Context(), id)
	if err != nil {
		h.log.Error(err.Error())
		response.JsonInternalServerError(w)
		return
	}
	if user == nil {
		response.JsonMessage(w, http.StatusNotFound, "user does not exist")
		return
	}
	response.JsonSuccess(w, struct {
		User UserResource `json:"user"`
	}{
		User: NewUserResource(user),
	})
}

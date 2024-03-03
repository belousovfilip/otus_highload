package signin

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"otus_highload/internal/lib/errs"
	"otus_highload/internal/lib/response"
)

type (
	request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	Auth interface {
		SignIn(ctx context.Context, email, password string) (string, error)
	}
)

type Handler struct {
	logger   *slog.Logger
	response *response.JSON
	auth     Auth
}

func New(logger *slog.Logger, response *response.JSON, auth Auth) *Handler {
	logger = logger.With("handler", "signin")
	return &Handler{logger: logger, response: response, auth: auth}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var req *request
	if req, err = h.getRequestData(r); err != nil {
		h.logger.With("handler", "getRequestData").Error(err.Error())
		h.response.InternalServerError(w)
		return
	}
	if err := h.validateRequest(req); err != nil {
		h.response.Message(w, http.StatusBadRequest, err.Error())
		return
	}
	var token string
	token, err = h.auth.SignIn(r.Context(), req.Email, req.Password)
	if err != nil {
		h.logger.With("Auth", "SignIn").Error(err.Error())
		h.response.InternalServerError(w)
		return
	}
	if errors.As(err, &errs.ErrUserByEmailNotFound{}) {
		h.response.Message(w, http.StatusBadRequest, "Credentials do not match")
		return
	}
	if errors.Is(err, errs.ErrUserPasswordIsNotEqual) {
		h.response.Message(w, http.StatusBadRequest, "Credentials do not match")
		return
	}
	h.response.Success(w, struct {
		Token string `json:"token"`
	}{Token: token})
}

func (h Handler) getRequestData(r *http.Request) (*request, error) {
	req := &request{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil && !errors.Is(io.EOF, err) {
		return nil, err
	}
	return req, nil
}

func (h Handler) validateRequest(r *request) error {
	validate := validator.New()
	if err := validate.Var(r.Email, "required,email,max=255"); err != nil {
		return errs.ErrInvalidEmail
	}
	if err := validate.Var(r.Password, "required,min=6,max=100"); err != nil {
		return errs.ErrInvalidPassword
	}
	return nil
}

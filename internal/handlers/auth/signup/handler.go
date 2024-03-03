package signup

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"otus_highload/internal/domain"
	"otus_highload/internal/lib/errs"
	"otus_highload/internal/lib/response"
	"time"

	"github.com/go-playground/validator/v10"
)

type (
	Auth interface {
		SignUp(context.Context, *domain.User) error
	}
)

type request struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (r request) toUser() *domain.User {
	return &domain.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Email:     r.Email,
		Password:  r.Password,
		CreatedAt: time.Time{},
	}
}

type Handler struct {
	logger   *slog.Logger
	response *response.JSON
	auth     Auth
}

func New(logger *slog.Logger, response *response.JSON, auth Auth) *Handler {
	logger = logger.With("handler", "signup")
	return &Handler{logger: logger, response: response, auth: auth}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	var req *request
	req, err = h.getRequestData(r)
	user := req.toUser()
	if err = h.auth.SignUp(r.Context(), user); err != nil {
		errUserExists := errs.ErrUserWithEmailAlreadyExists{}
		if errors.As(err, &errUserExists) {
			h.response.ErrorW(w, http.StatusBadRequest, errUserExists)
			return
		}
		h.logger.With("auth", "SignUp").Error(err.Error())
		h.response.InternalServerError(w)
		return
	}
	h.response.Message(w, http.StatusCreated, http.StatusText(http.StatusCreated))
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
	if err := validate.Var(r.FirstName, "required,min=3,max=255"); err != nil {
		return errs.ErrInvalidFirstName
	}
	if err := validate.Var(r.LastName, "required,min=3,max=255"); err != nil {
		return errs.ErrInvalidLastName
	}
	return nil
}

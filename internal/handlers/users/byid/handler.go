package byid

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"otus_highload/internal/domain"
	"otus_highload/internal/lib/errs"
	"otus_highload/internal/lib/response"
	"strconv"
	"time"
)

type (
	UsersReader interface {
		UserByID(context.Context, int64) (*domain.User, error)
	}
)

type Handler struct {
	logger      *slog.Logger
	response    *response.JSON
	usersReader UsersReader
}

func NewHandler(
	logger *slog.Logger,
	response *response.JSON,
	usersReader UsersReader,
) *Handler {
	logger = logger.With("handler", "UserByID")
	return &Handler{
		logger:      logger,
		response:    response,
		usersReader: usersReader,
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		h.response.Message(w, http.StatusBadRequest, "invalid user id")
		return
	}
	user, err := h.usersReader.UserByID(r.Context(), id)
	if err != nil {
		errUserNotFound := errs.ErrUserByIDNotFound{}
		if errors.As(err, &errUserNotFound) {
			h.response.ErrorW(w, http.StatusNotFound, errUserNotFound)
			return
		}
		h.logger.Error(err.Error())
		h.response.InternalServerError(w)
		return
	}
	h.response.Success(w, struct {
		User UserResource `json:"user"`
	}{
		User: NewUserResource(user),
	})
}

type UserResource struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Age       int       `json:"age,omitempty"`
	Gender    string    `json:"gender"`
	City      string    `json:"city"`
	Interests string    `json:"interests"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewUserResource(u *domain.User) UserResource {
	return UserResource{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Age:       u.Age,
		Gender:    u.Gender,
		City:      u.City,
		Interests: u.Interests,
		CreatedAt: u.CreatedAt,
	}
}

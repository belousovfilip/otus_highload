package signup

import (
	"otus_highload/internal/lib/errs"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	h := Handler{}

	longPassword := "password_password_password_password"
	longPassword += longPassword + longPassword + longPassword
	cases := []struct {
		name string
		req  *request
		out  error
	}{
		{
			name: "invalid email",
			req: &request{
				Email:    "testtest.com",
				Password: "password",
			},
			out: errs.ErrInvalidEmail,
		},
		{
			name: "invalid password: short",
			req: &request{
				Email:    "test@gmain.com",
				Password: "pass",
			}, out: errs.ErrInvalidPassword,
		},
		{
			name: "invalid password: long",
			req: &request{
				Email:    "test@gmain.com",
				Password: longPassword,
			},
			out: errs.ErrInvalidPassword,
		},
		{
			name: "invalid first name: short",
			req: &request{
				Email:     "test@gmain.com",
				Password:  "password",
				FirstName: "ss",
				LastName:  "sssssss",
			},
			out: errs.ErrInvalidFirstName,
		},
		{
			name: "invalid last name: short",
			req: &request{
				Email:     "test@gmain.com",
				Password:  "password",
				FirstName: "ssssssss",
				LastName:  "ss",
			},
			out: errs.ErrInvalidLastName,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := h.validateRequest(tt.req)
			require.ErrorIs(t, err, tt.out)
		})
	}
}

package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"otus_highload/internal/models"
)

func (s UserStorage) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	u := &models.User{}
	q := `
		SELECT id, email, password, created_at,
		       COALESCE(first_name, '') as first_name, 
		       COALESCE(last_name, '') as  last_name
    	FROM users  
    	WHERE email=?
`
	err := s.db.GetContext(ctx, u, q, email)
	if err == nil {
		return u, nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return nil, fmt.Errorf("user storage: fiend user by email; %s", err)
}

package users

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"otus_highload/internal/models"
)

func (s UserStorage) FindUserById(ctx context.Context, id int64) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT 
		   id, email, password, created_at,
		   COALESCE(first_name, '') as first_name, 
		   COALESCE(last_name, '') as last_name,
		   COALESCE(age, 0) as age,
		   COALESCE(gender, '') as gender,
		   COALESCE(city, '') as city,
		   COALESCE(interests, '') as interests
    	FROM users  
    	WHERE id=?
`
	err := s.db.GetContext(ctx, user, query, id)
	if err == nil {
		return user, nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	return nil, fmt.Errorf("user storage: find user by id; %s", err.Error())
}

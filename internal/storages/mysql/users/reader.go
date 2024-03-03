package users

import (
	"context"
	"errors"
	"fmt"
	"otus_highload/internal/domain"
	"otus_highload/internal/lib/db"
	"otus_highload/internal/lib/errs"

	"github.com/jackc/pgx/v5"
)

type Reader struct {
	conn *db.Pool
}

func NewReader(conn *db.Pool) *Reader {
	return &Reader{
		conn: conn,
	}
}

func (s Reader) UserByEmail(ctx context.Context, email string) (*domain.User, error) {
	conn := s.conn.GetConn()
	user := &domain.User{}
	const query = `
		SELECT id, email, password, created_at,
			   COALESCE(first_name, '') as first_name, 
			   COALESCE(last_name, '') as  last_name
		FROM users  
		WHERE email=$1
	`
	row := conn.QueryRow(ctx, query, email)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.FirstName,
		&user.LastName,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errs.ErrUserByEmailNotFound{Email: email}
	}
	if err != nil {
		return nil, fmt.Errorf("storage:users:reader:UserByEmail: %w", err)
	}
	return user, nil
}

func (s Reader) UserByID(ctx context.Context, id int64) (*domain.User, error) {
	conn := s.conn.GetConn()
	user := &domain.User{}
	const query = `
		SELECT 
		   id, email, password, created_at,
		   COALESCE(first_name, '') as first_name, 
		   COALESCE(last_name, '') as last_name,
		   COALESCE(age, 0) as age,
		   COALESCE(gender, '') as gender,
		   COALESCE(city, '') as city,
		   COALESCE(interests, '') as interests
    	FROM users  
    	WHERE id=$1
	`
	row := conn.QueryRow(ctx, query, id)
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.FirstName,
		&user.LastName,
		&user.Age,
		&user.Gender,
		&user.City,
		&user.Interests,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errs.ErrUserByIDNotFound{ID: id}
	}
	if err != nil {
		return nil, fmt.Errorf("storage:users:reader:UserByID: %w", err)
	}
	return user, nil
}

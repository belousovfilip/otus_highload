package users

import (
	"context"
	"fmt"
	"otus_highload/internal/domain"
	"otus_highload/internal/lib/db"
)

type Writer struct {
	conn *db.Pool
}

func NewWriter(conn *db.Pool) *Writer {
	return &Writer{conn: conn}
}

func (s Writer) Add(ctx context.Context, u *domain.User) error {
	conn := s.conn.GetConn()
	const query = "INSERT INTO users (email, password, created_at) VALUES ($1, $2, $3)"
	_, err := conn.Exec(ctx, query, u.Email, u.Password, u.CreatedAt)
	if err != nil {
		return fmt.Errorf("storage:users:writer:Add: %w", err)
	}
	return nil
}

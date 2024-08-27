package entity

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func InitDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgresql:postgres:postgres@localhost/bogolang?sslmode=disable")
	if err != nil {
	}
	defer conn.Close(context.Background())
	return conn
}

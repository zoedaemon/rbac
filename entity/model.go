package entity

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func InitDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgresql:postgres:postgres@localhost/bogolang?sslmode=disable")
	if err != nil {
		log.Fatal("Error Connecting to DB: ", err)
		return nil
	}
	return conn
}

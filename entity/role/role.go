package user

import "time"

type Role struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

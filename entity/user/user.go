package user

import (
	api "rbac/proto"
	"time"
)

type User struct {
	api.UnimplementedAPIServer

	ID         int64     `json:"id"`
	RoleID     int64     `json:"role_id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Name       string    `json:"name"`
	LastAccess time.Time `json:"last_access"`
	CreatedAt  time.Time `json:"created_at"`
	UpdateAt   time.Time `json:"update_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	DB         interface{}
}

func (user *User) SetDB(db interface{}) {
	user.DB = db
}

package user

import (
	"rbac/entity/user"
)

type User interface {
	GetAllUser() ([]user.User, error)
}

// func GetAllUser() {

// }

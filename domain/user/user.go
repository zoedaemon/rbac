package user

import (
	"rbac/entity"
)

type User interface {
	GetAllUser() ([]entity.User, error)
}

// func GetAllUser() {

// }

package entity

type User struct {
	ID     int64  `json:"id"`
	RoleID int64  `json:"role_id"`
	Email  string `json:"email"`
}

package user

type Role struct {
	ID      int64  `json:"id"`
	RoleID  int64  `json:"role_id"`
	Route   string `json:"route"`
	Section string `json:"section"`
	Path    string `json:"path"`
}

package user

import (
	"context"
	api "rbac/proto"
	// "google.golang.org/grpc"
)

type UserRequest struct {
	Id string
}

func (userEntity *User) GetAllUser(ctx context.Context, req *api.UserId) (*api.User, error) {
	user := &api.User{}
	user.Id = uint64(userEntity.ID)
	user.Email = userEntity.Email
	return user, nil
}

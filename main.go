package main

import (
	"context"
	"log"
	"net"

	// user_domain "rbac/domain/user"
	entity "rbac/entity"
	user "rbac/entity/user"
	api "rbac/proto"

	"google.golang.org/grpc"
)

func main() {
	conn := entity.InitDB()
	defer conn.Close(context.Background())
	// var UserDomain *user_domain.User
	userEntity := &user.User{}
	userEntity.SetDB(conn)
	s := grpc.NewServer()
	api.RegisterAPIServer(s, userEntity)
	lis, err := net.Listen("tcp", "5555")
	if err != nil {
		log.Fatal("Failed to Listen: ", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to Serve: ", err)
	}
}

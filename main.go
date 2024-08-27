package main

import (
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
	// var UserDomain *user_domain.User
	userEntity := &user.User{}
	userEntity.SetDB(conn)
	s := grpc.NewServer()
	api.RegisterAPIServer(s, userEntity)
	lis, err := net.Listen("tcp", "5555")
	if err != nil {
		log.Fatal("Failed to Listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to Serve: %v", err)
	}
}

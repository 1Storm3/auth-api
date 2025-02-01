package srv

import (
	"user-api/internal/delivery"
	"user-api/pkg/proto/gengrpc"
)

type Server struct {
	gengrpc.UnimplementedUserServiceServer
	gengrpc.UnimplementedAuthServiceServer
	userService delivery.UserService
	authService delivery.AuthService
}

func NewGRPCServer(
	userService delivery.UserService,
	authService delivery.AuthService) *Server {
	return &Server{
		userService: userService,
		authService: authService,
	}
}

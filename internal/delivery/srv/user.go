package srv

import (
	"context"
	"log"

	"github.com/google/uuid"
	"user-api/internal/converter"
	"user-api/internal/dto"
	"user-api/pkg/proto/gengrpc"
)

func (s *Server) Create(ctx context.Context, req *gengrpc.CreateUserRequest) (*gengrpc.CreateUserResponse, error) {
	if err := req.Validate(); err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := s.userService.Create(ctx, dto.CreateUserDto{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})

	if err != nil {
		return nil, err
	}

	return &gengrpc.CreateUserResponse{Data: converter.DtoUserToGrpc(result)}, nil
}

func (s *Server) GetOne(ctx context.Context, req *gengrpc.GetOneUserRequest) (*gengrpc.GetOneUserResponse, error) {
	if err := req.Validate(); err != nil {
		log.Println(err)
		return nil, err
	}

	userID := req.GetId()
	result, err := s.userService.GetOne(ctx, userID)

	if err != nil {
		return nil, err
	}

	return &gengrpc.GetOneUserResponse{Data: converter.DtoUserToGrpc(result)}, nil
}

func (s *Server) Update(ctx context.Context, req *gengrpc.UpdateUserRequest) (*gengrpc.UpdateUserResponse, error) {
	if err := req.Validate(); err != nil {
		log.Println(err)
		return nil, err
	}

	userId, err := uuid.Parse(req.GetId())

	if err != nil {
		return nil, err
	}

	result, err := s.userService.Update(ctx, dto.UpdateUserDto{
		ID:        userId,
		Email:     req.GetEmail(),
		FirstName: req.GetFirstName(),
	})

	if err != nil {
		return nil, err
	}

	return &gengrpc.UpdateUserResponse{Data: converter.DtoUserToGrpc(result)}, nil
}

func (s *Server) Delete(ctx context.Context, req *gengrpc.DeleteUserRequest) (*gengrpc.DeleteUserResponse, error) {
	if err := req.Validate(); err != nil {
		log.Println(err)
		return nil, err
	}

	userID := req.GetId()
	err := s.userService.Delete(ctx, userID)

	if err != nil {
		return nil, err
	}

	return &gengrpc.DeleteUserResponse{
		Data: "Пользователь удален",
	}, nil
}

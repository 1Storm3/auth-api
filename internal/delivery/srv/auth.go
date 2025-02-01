package srv

import (
	"context"
	"log"

	"github.com/google/uuid"
	"user-api/internal/converter"
	"user-api/internal/dto"
	"user-api/pkg/proto/gengrpc"
)

func (s *Server) Login(ctx context.Context, req *gengrpc.LoginRequest) (*gengrpc.LoginResponse, error) {
	if err := req.Validate(); err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := s.authService.Login(ctx, dto.LoginDto{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})

	if err != nil {
		return nil, err
	}

	return &gengrpc.LoginResponse{
		Token: result.AccessToken,
	}, nil
}

func (s *Server) Register(ctx context.Context, req *gengrpc.RegisterRequest) (*gengrpc.RegisterResponse, error) {
	if err := req.Validate(); err != nil {
		log.Println(err)
		return nil, err
	}

	result, err := s.authService.Register(ctx, dto.RegisterDto{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})

	if err != nil {
		return nil, err
	}

	return &gengrpc.RegisterResponse{
		Token: result.AccessToken,
	}, nil
}

func (s *Server) Verify(ctx context.Context, req *gengrpc.VerifyRequest) (*gengrpc.VerifyResponse, error) {
	if err := req.Validate(); err != nil {
		log.Println(err)
		return nil, err
	}

	return &gengrpc.VerifyResponse{}, nil
}

func (s *Server) Me(ctx context.Context, req *gengrpc.MeRequest) (*gengrpc.MeResponse, error) {
	if err := req.Validate(); err != nil {
		log.Println(err)
		return nil, err
	}

	userID, err := uuid.Parse(req.GetUserId())
	if err != nil {
		log.Println("Invalid UUID format:", err)
		return nil, err
	}

	result, err := s.authService.Me(ctx, dto.MeDto{
		UserID: userID,
	})

	if err != nil {
		return nil, err
	}

	return &gengrpc.MeResponse{
		User: converter.MeResponseDtoToAuthUser(result),
	}, nil
}

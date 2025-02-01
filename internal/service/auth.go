package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"user-api/internal/config"
	"user-api/internal/converter"
	"user-api/internal/delivery"
	"user-api/internal/dto"
	"user-api/pkg/proto/gengrpc"
)

type AuthService struct {
	gengrpc.UnimplementedAuthServiceServer
	userService  delivery.UserService
	tokenService delivery.TokenService
	cfg          *config.Config
}

func NewAuthService(
	userService delivery.UserService,
	tokenService delivery.TokenService,
	cfg *config.Config) *AuthService {
	return &AuthService{
		userService:  userService,
		tokenService: tokenService,
		cfg:          cfg,
	}
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginDto) (dto.LoginResponseDto, error) {
	user, err := s.userService.GetOneByEmail(ctx, req.Email)

	if err != nil || !s.userService.CheckPassword(ctx, user, req.Password) {
		log.Println(err)
		return dto.LoginResponseDto{}, errors.New("email or password is incorrect")
	}
	jwtKey := []byte(s.cfg.App.JwtSecretKey)

	expiresIn, err := time.ParseDuration(s.cfg.App.JwtExpiresIn)

	if err != nil {
		return dto.LoginResponseDto{}, err
	}

	tokenString, err := s.tokenService.GenerateToken(jwtKey, user.ID.String(), user.Role, expiresIn)

	if err != nil {
		return dto.LoginResponseDto{}, err
	}

	return dto.LoginResponseDto{
		AccessToken: tokenString,
	}, nil
}

func (s *AuthService) Me(ctx context.Context, req dto.MeDto) (dto.MeResponseDto, error) {
	user, err := s.userService.GetOne(ctx, req.UserID.String())

	if err != nil {
		return dto.MeResponseDto{}, err
	}

	if user.ID == uuid.Nil {
		return dto.MeResponseDto{}, errors.New("user not found")
	}

	return converter.UserResponseDtoToMeResponseDto(user), nil
}

func (s *AuthService) Register(ctx context.Context, req dto.RegisterDto) (dto.RegisterResponseDto, error) {
	userExist, err := s.userService.GetOneByEmail(ctx, req.Email)

	if err != nil {
		return dto.RegisterResponseDto{}, err
	}

	if userExist.ID != uuid.Nil {
		return dto.RegisterResponseDto{}, errors.New("user already exist")
	}

	hash, err := s.userService.HashPassword(ctx, req.Password)

	if err != nil {
		return dto.RegisterResponseDto{}, err
	}

	user, err := s.userService.Create(ctx, dto.CreateUserDto{
		Email:    req.Email,
		Password: hash,
	})

	if err != nil {
		return dto.RegisterResponseDto{}, err
	}

	jwtKey := []byte(s.cfg.App.JwtSecretKey)

	expiresIn, err := time.ParseDuration(s.cfg.App.JwtExpiresIn)

	if err != nil {
		return dto.RegisterResponseDto{}, err
	}

	token, err := s.tokenService.GenerateToken(jwtKey, user.ID.String(), user.Role, expiresIn)

	if err != nil {
		return dto.RegisterResponseDto{}, err
	}

	return dto.RegisterResponseDto{
		AccessToken: token,
	}, nil
}

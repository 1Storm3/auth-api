package delivery

import (
	"context"
	"time"

	"user-api/internal/dto"
	"user-api/internal/model"
	"user-api/pkg/proto/gengrpc"
)

type UserService interface {
	Create(ctx context.Context, req dto.CreateUserDto) (dto.ResponseUserDto, error)
	Update(ctx context.Context, req dto.UpdateUserDto) (dto.ResponseUserDto, error)
	Delete(ctx context.Context, id string) error
	GetOne(ctx context.Context, id string) (dto.ResponseUserDto, error)
	GetOneByEmail(ctx context.Context, email string) (model.User, error)
	CheckPassword(ctx context.Context, user model.User, password string) bool
	HashPassword(ctx context.Context, password string) (string, error)
}

type AuthService interface {
	Login(ctx context.Context, req dto.LoginDto) (dto.LoginResponseDto, error)
	Register(ctx context.Context, req dto.RegisterDto) (dto.RegisterResponseDto, error)
	Verify(ctx context.Context, req *gengrpc.VerifyRequest) (*gengrpc.VerifyResponse, error)
	Me(ctx context.Context, req dto.MeDto) (dto.MeResponseDto, error)
}

type TokenService interface {
	GenerateToken(jwtKey []byte, userID, role string, duration time.Duration) (string, error)
}

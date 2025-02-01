package service

import (
	"context"

	"golang.org/x/crypto/bcrypt"
	"user-api/internal/converter"
	"user-api/internal/dto"
	"user-api/internal/model"
	"user-api/pkg/proto/gengrpc"
)

type UserService struct {
	gengrpc.UnimplementedUserServiceServer
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) Create(
	ctx context.Context,
	req dto.CreateUserDto,
) (dto.ResponseUserDto, error) {
	result, err := u.userRepo.Create(ctx, model.User{
		Email:        req.Email,
		PasswordHash: req.Password,
	})
	if err != nil {
		return dto.ResponseUserDto{}, err
	}

	return converter.ModelUserToDto(result), nil
}

func (u *UserService) GetOne(ctx context.Context, id string) (dto.ResponseUserDto, error) {
	result, err := u.userRepo.GetOne(ctx, id)
	if err != nil {
		return dto.ResponseUserDto{}, err
	}

	return converter.ModelUserToDto(result), nil
}

func (u *UserService) Update(ctx context.Context, req dto.UpdateUserDto) (dto.ResponseUserDto, error) {
	result, err := u.userRepo.Update(ctx, model.User{
		ID:        req.ID,
		Email:     req.Email,
		FirstName: &req.FirstName,
	})
	if err != nil {
		return dto.ResponseUserDto{}, err
	}

	return converter.ModelUserToDto(result), nil
}

func (u *UserService) Delete(ctx context.Context, id string) error {
	return u.userRepo.Delete(ctx, id)
}

func (u *UserService) GetOneByEmail(ctx context.Context, email string) (model.User, error) {
	result, err := u.userRepo.GetOneByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}

	return result, nil
}

func (u *UserService) CheckPassword(_ context.Context, user model.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	return err == nil
}

func (u *UserService) HashPassword(_ context.Context, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

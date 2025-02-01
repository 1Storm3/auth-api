package service

import (
	"context"

	"user-api/internal/model"
)

type UserRepo interface {
	Create(ctx context.Context, req model.User) (model.User, error)
	Update(ctx context.Context, req model.User) (model.User, error)
	Delete(ctx context.Context, userID string) error
	GetOne(ctx context.Context, userID string) (model.User, error)
	GetOneByEmail(ctx context.Context, email string) (model.User, error)
}

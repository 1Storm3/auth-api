package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"user-api/internal/converter"
	"user-api/internal/model"
	"user-api/pkg/gensqlc"
)

type UserRepo struct {
	pool    *pgxpool.Pool
	queries *gensqlc.Queries
}

func NewUserRepo(pool *pgxpool.Pool, queries *gensqlc.Queries) *UserRepo {
	return &UserRepo{
		pool:    pool,
		queries: queries,
	}
}

func (u *UserRepo) Create(ctx context.Context, req model.User) (model.User, error) {
	result, err := u.queries.CreateUser(ctx, gensqlc.CreateUserParams{
		PasswordHash: req.PasswordHash,
		Email:        req.Email,
	})
	if err != nil {
		return model.User{}, err
	}

	return converter.SqlcUserToModel(result), nil
}

func (u *UserRepo) GetOneByEmail(ctx context.Context, email string) (model.User, error) {
	result, err := u.queries.GetOneByEmail(ctx, gensqlc.GetOneByEmailParams{
		Email: email,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, nil
		}
		return model.User{}, err
	}

	return converter.SqlcUserToModel(result), nil
}

func (u *UserRepo) Update(ctx context.Context, req model.User) (model.User, error) {
	result, err := u.queries.UpdateUser(ctx, gensqlc.UpdateUserParams{
		ID:        req.ID,
		Email:     req.Email,
		FirstName: req.FirstName,
	})
	if err != nil {
		return model.User{}, err
	}

	return converter.SqlcUserToModel(result), nil
}

func (u *UserRepo) Delete(ctx context.Context, id string) error {
	userId, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	err = u.queries.DeleteUser(ctx, gensqlc.DeleteUserParams{
		ID: userId,
	})

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) GetOne(ctx context.Context, id string) (model.User, error) {
	userId, err := uuid.Parse(id)
	if err != nil {
		return model.User{}, err
	}

	result, err := u.queries.GetUser(ctx, gensqlc.GetUserParams{
		ID: userId,
	})

	if err != nil {
		return model.User{}, err
	}

	return converter.SqlcUserToModel(result), nil
}

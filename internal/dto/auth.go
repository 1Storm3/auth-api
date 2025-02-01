package dto

import (
	"github.com/google/uuid"
)

type LoginDto struct {
	Email    string
	Password string
}

type LoginResponseDto struct {
	AccessToken string
}

type MeDto struct {
	UserID uuid.UUID
}

type MeResponseDto struct {
	ID           uuid.UUID
	FirstName    *string
	LastName     *string
	MiddleName   *string
	Email        string
	Role         string
	Status       string
	Phone        *string
	Address      *string
	LastActivity string
	IsVerified   bool
	Photo        *string
	CreatedAt    string
	UpdatedAt    string
}

type RegisterDto struct {
	Email    string
	Password string
}

type RegisterResponseDto struct {
	AccessToken string
}

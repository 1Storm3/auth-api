package dto

import (
	"github.com/google/uuid"
)

type CreateUserDto struct {
	Email    string
	Password string
}

type UpdateUserDto struct {
	ID            uuid.UUID
	FirstName     string
	LastName      string
	MiddleName    string
	Email         string
	Role          string
	Status        string
	Photo         string
	Phone         string
	Address       string
	VerifiedToken string
	IsVerified    bool
}
type ResponseUserDto struct {
	ID            uuid.UUID
	FirstName     *string
	LastName      *string
	MiddleName    *string
	Email         string
	Role          string
	Status        string
	Photo         *string
	Phone         *string
	Address       *string
	VerifiedToken *string
	LastActivity  string
	IsVerified    bool
	CreatedAt     string
	UpdatedAt     string
}

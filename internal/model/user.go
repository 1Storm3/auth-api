package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `db:"id"`
	FirstName     *string   `db:"first_name"`
	LastName      *string   `db:"last_name"`
	MiddleName    *string   `db:"middle_name"`
	Email         string    `db:"email"`
	PasswordHash  string    `db:"password_hash"`
	Role          string    `db:"role"`
	Status        string    `db:"status"`
	Photo         *string   `db:"photo"`
	Phone         *string   `db:"phone"`
	Address       *string   `db:"address"`
	VerifiedToken *string   `db:"verified_token"`
	LastActivity  string    `db:"lastActivity"`
	IsVerified    bool      `db:"is_verified"`
	CreatedAt     string    `db:"created_at"`
	UpdatedAt     string    `db:"updated_at"`
}

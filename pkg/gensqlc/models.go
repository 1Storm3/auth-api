// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package gensqlc

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID `json:"id"`
	FirstName     *string   `json:"first_name"`
	LastName      *string   `json:"last_name"`
	MiddleName    *string   `json:"middle_name"`
	Email         string    `json:"email"`
	PasswordHash  string    `json:"password_hash"`
	Role          string    `json:"role"`
	Status        string    `json:"status"`
	Photo         *string   `json:"photo"`
	Phone         *string   `json:"phone"`
	Address       *string   `json:"address"`
	VerifiedToken *string   `json:"verified_token"`
	LastActivity  string    `json:"last_activity"`
	IsVerified    bool      `json:"is_verified"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

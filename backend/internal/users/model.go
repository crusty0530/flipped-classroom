package users

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	Teacher Role = "TEACHER"
	Student Role = "STUDENT"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	DisplayName  string    `json:"display_name"`
	Role         Role      `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
}

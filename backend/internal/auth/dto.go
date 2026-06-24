package auth

import (
	"github.com/crusty0530/flipped-classroom/backend/internal/users"
)

type RegisterRequest struct {
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	DisplayName string     `json:"display_name"`
	Role        users.Role `json:"role"`
}

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

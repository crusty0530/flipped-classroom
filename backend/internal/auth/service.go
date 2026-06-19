package auth

import (
	"errors"

	"github.com/crusty0530/flipped-classroom/backend/internal/users"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userService *users.Service
}

func NewService(userService *users.Service) *Service {
	return &Service{userService: userService}
}

// TODO: i had to go to work, will finish register and login functions
func (s *Service) Register(request RegisterRequest) error {
	user_user, user_err := s.userService.FindByUserOrEmail(request.Username)
	email_user, email_err := s.userService.FindByUserOrEmail(request.Email)

	if user_user != nil {
		return errors.New("Username already exists!")
	}
	if email_user != nil {
		return errors.New("Email already exists!")
	}
	if user_err != nil {
		return user_err
	}
	if email_err != nil {
		return email_err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	user := users.User{}
}

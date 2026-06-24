package auth

import (
	"errors"
	"os"
	"time"

	"github.com/crusty0530/flipped-classroom/backend/internal/users"
	"github.com/google/uuid"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	userService *users.Service
}

func NewService(userService *users.Service) *Service {
	return &Service{userService: userService}
}

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

	user := users.User{
		Username:     request.Username,
		Email:        request.Email,
		PasswordHash: string(hashedPassword),
		DisplayName:  request.DisplayName,
		Role:         request.Role,
	}

	err = s.userService.InsertUser(&user)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Login(request LoginRequest) (*users.User, error) {
	user, err := s.userService.FindByUserOrEmail(request.UsernameOrEmail)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("Invalid Credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password))

	if err != nil {
		return nil, errors.New("Invalid Credentials")
	}

	return user, nil
}

func (s *Service) GenerateJWTToken(id uuid.UUID, role users.Role) (string, error) {
	claims := jwt.MapClaims{
		"id":   id,
		"role": role,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
		"iat":  time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	request := RegisterRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.Register(request)

	if err != nil {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprint(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	request := LoginRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.service.Login(request)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, err)
		return
	}

	token, err := h.service.GenerateJWTToken(user.ID, user.Role)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	secure := os.Getenv("APP_ENV") == "production"

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",                              // Scope of the cookie
		Expires:  time.Now().Add(15 * time.Minute), // Expiration time
		HttpOnly: true,                             // Prevents JavaScript access
		Secure:   secure,                           // Requires HTTPS
		SameSite: http.SameSiteLaxMode,             // CSRF protection
	}

	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

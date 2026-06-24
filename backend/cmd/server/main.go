package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crusty0530/flipped-classroom/backend/internal/auth"
	"github.com/crusty0530/flipped-classroom/backend/internal/db"
	"github.com/crusty0530/flipped-classroom/backend/internal/users"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	database := db.Connect()
	userRepo := users.NewRepository(database)
	userService := users.NewService(userRepo)
	authService := auth.NewService(userService)
	authHandler := auth.NewHandler(authService)

	defer database.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	})

	mux.HandleFunc("/api/auth/register", authHandler.Register)
	mux.HandleFunc("/api/auth/login", authHandler.Login)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

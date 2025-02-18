package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"nimblestack/database"
	"nimblestack/views"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	Queries   *database.Queries
	jwtSecret []byte
	tokenExp  time.Duration
}

func NewUserHandler(queries *database.Queries, jwtSecret []byte) *UserHandler {
	return &UserHandler{
		Queries:   queries,
		jwtSecret: jwtSecret,
		tokenExp:  24 * time.Hour,
	}
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := views.Auth().Render(r.Context(), w); err != nil {
			log.Println("Error rendering view:", err)
		}
	} else {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			log.Printf("form error: %v\n", err)
			return

		}
		req := loginRequest{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		user, err := h.Queries.GetUser(ctx, req.Email)
		if err != nil {
			if !errors.Is(err, sql.ErrNoRows) {
				log.Printf("Database error: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":    user.ID,
			"email": user.Email,
			"exp":   time.Now().Add(h.tokenExp).Unix(),
		})

		tokenString, err := token.SignedString(h.jwtSecret)
		if err != nil {
			log.Printf("Token signing error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"token":   tokenString,
			"message": "Login successful",
			"email":   user.Email,
		})

	}

}

// RegisterUser
func (h *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if err := views.Auth().Render(r.Context(), w); err != nil {
			log.Println("Error rendering view:", err)
		}
	} else {

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			log.Printf("form error: %v\n", err)
			return
		}
		req := database.CreateUserParams{
			Username:  r.FormValue("Username"),
			Email:     r.FormValue("email"),
			Firstname: r.FormValue("firstname"),
			Lastname:  r.FormValue("lastname"),
			Password:  r.FormValue("password"),
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Password hashing error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		req.Password = string(hashedPassword)

		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		_, err = h.Queries.GetUser(ctx, req.Email)
		if err == nil {
			http.Error(w, "Email already registered", http.StatusConflict)
			return
		} else if !errors.Is(err, sql.ErrNoRows) {
			log.Printf("Database error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if _, err := h.Queries.CreateUser(ctx, req); err != nil {
			log.Printf("User creation error: %v", err)
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": req.Email,
			"exp":   time.Now().Add(h.tokenExp).Unix(),
		})

		tokenString, err := token.SignedString(h.jwtSecret)
		if err != nil {
			log.Printf("Token signing error: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"token":   tokenString,
			"message": "Login successful",
			"email":   req.Email,
		})
	}

}

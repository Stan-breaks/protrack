package apis

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"nimblestack/database"
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
		tokenExp:  time.Hour,
	}
}

type loginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid http method", http.StatusMethodNotAllowed)
		return
	}
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
	http.SetCookie(w, &http.Cookie{
		Name:     "auth_token",
		Value:    tokenString,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(h.tokenExp),
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
	w.WriteHeader(http.StatusOK)
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

// Define a key type for context values to avoid collisions
type contextKey string

// Define the key for JWT claims
const jwtClaimsKey contextKey = "jwtClaims"

// User represents a user in our system
type User struct {
	Username string
	Role     string
}

// JWTMiddleware verifies the token and adds claims to the context
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Check if the Authorization header has the correct format
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		// Parse and validate the token
		tokenString := bearerToken[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Return your secret key here
			return []byte("your-secret-key"), nil
		})

		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Check if the token is valid
		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		// Add claims to the request context using the typed key
		ctx := context.WithValue(r.Context(), jwtClaimsKey, claims)

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Example handler that uses the JWT claims
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// Extract claims from context using the typed key
	claims, ok := r.Context().Value(jwtClaimsKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Use the claims
	username, ok := claims["username"].(string)
	if !ok {
		http.Error(w, "Invalid username claim", http.StatusBadRequest)
		return
	}

	role, _ := claims["role"].(string) // Role is optional

	// Respond with the user info
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Hello, %s!", "role": "%s"}`, username, role)
}

// Function to generate a JWT token for a user
func generateToken(user User) (string, error) {
	// Create the claims
	claims := jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"iat":      time.Now().Unix(),                     // Issued at time
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with our secret key
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Login handler to generate and return a token
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// In a real application, you would validate credentials here
	// For this example, we'll just create a token for a demo user

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Create a user (in a real app, you'd verify credentials)
	user := User{
		Username: "demo_user",
		Role:     "admin",
	}

	// Generate a token
	token, err := generateToken(user)
	if err != nil {
		http.Error(w, "Error generating token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the token
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"token": "%s"}`, token)
}

func main() {
	// Create a new mux
	mux := http.NewServeMux()

	// Login route (unprotected)
	mux.HandleFunc("/login", loginHandler)

	// Protected route with JWT middleware
	protectedRoute := JWTMiddleware(http.HandlerFunc(protectedHandler))
	mux.Handle("/protected", protectedRoute)

	// Start the server
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

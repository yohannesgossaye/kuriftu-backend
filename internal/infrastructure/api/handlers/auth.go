package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/yohannesgossaye/kuriftu-backend/internal/application/auth"
)

// RegisterHandler godoc
// @Summary Register a new user
// @Description Creates a new user account with personal information
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "User registration details"
// @Success 201 {object} RegisterResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/register [post]
func RegisterHandler(svc *auth.Service, log *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error().Err(err).Msg("Invalid request body")
			http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
			return
		}

		user, err := svc.Register(req.FirstName, req.LastName, req.Email, req.Password, req.Phone, req.UserType)
		if err != nil {
			log.Error().Err(err).Msg("Registration failed")
			http.Error(w, `{"error": "Registration failed: `+err.Error()+`"}`, http.StatusBadRequest)
			return
		}

		resp := RegisterResponse{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.Phone,
			UserType:  user.UserType,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			IsActive:  user.IsActive,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Error().Err(err).Msg("Failed to encode response")
			http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
			return
		}
	}
}

// LoginHandler godoc
// @Summary Login a user
// @Description Authenticates a user and returns a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} LoginResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/login [post]
func LoginHandler(svc *auth.Service, log *zerolog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error().Err(err).Msg("Invalid request body")
			http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
			return
		}

		token, err := svc.Login(req.Email, req.Password)
		if err != nil {
			log.Error().Err(err).Msg("Login failed")
			http.Error(w, `{"error": "Login failed: `+err.Error()+`"}`, http.StatusUnauthorized)
			return
		}

		resp := LoginResponse{Token: token}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Error().Err(err).Msg("Failed to encode response")
			http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
			return
		}
	}
}

type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	UserType  string `json:"user_type"`
}

type RegisterResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	UserType  string `json:"user_type"`
	CreatedAt string `json:"created_at"`
	IsActive  bool   `json:"is_active"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
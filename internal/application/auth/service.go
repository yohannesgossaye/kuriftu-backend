package auth

import (
	"database/sql"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/yohannesgossaye/kuriftu-backend/internal/config"
	"github.com/yohannesgossaye/kuriftu-backend/internal/domain/auth"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo   auth.Repository
	config *config.Config
}

func NewService(repo auth.Repository, cfg *config.Config) *Service {
	return &Service{repo: repo, config: cfg}
}

func (s *Service) Register(firstName, lastName, email, password, phone, userType string) (auth.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return auth.User{}, err
	}

	user := auth.User{
		FirstName:    firstName,
		LastName:     lastName,
		Email:        email,
		PasswordHash: string(hash),
		Phone:        phone,
		UserType:     userType,
		IsActive:     true,
	}

	createdUser, err := s.repo.CreateUser(user)
	if err != nil {
		return auth.User{}, err
	}

	return createdUser, nil
}

func (s *Service) Login(email, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("email and password are required")
	}

	user, err := s.repo.GetUserByEmail(email)
	if err == sql.ErrNoRows {
		return "", errors.New("user not found")
	}
	if err != nil {
		return "", err
	}
	if !user.IsActive {
		return "", errors.New("user is inactive")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"email":     user.Email,
		"user_type": user.UserType,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.config.JWTSecret))
	if err != nil {
		return "", err
	}

	if err := s.repo.UpdateLastLoginAt(user.ID); err != nil {
		return "", err
	}

	return tokenString, nil
}
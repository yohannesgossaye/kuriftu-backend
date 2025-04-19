package auth

import (
	"github.com/yohannesgossaye/kuriftu-backend/internal/domain/auth"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo auth.Repository
}

func NewService(repo auth.Repository) *Service {
	return &Service{repo: repo}
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

package auth

import (
	"context"
	"database/sql"

	"github.com/yohannesgossaye/kuriftu-backend/internal/domain/auth"
	"github.com/yohannesgossaye/kuriftu-backend/internal/infrastructure/database/sqlc"
)

type Repository struct {
	db *sqlc.Queries
}

func NewRepository(db *sqlc.Queries) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateUser(user auth.User) (auth.User, error) {
	created, err := r.db.CreateUser(context.Background(), sqlc.CreateUserParams{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Phone:        user.Phone,
		UserType:     user.UserType,
	})
	if err != nil {
		return auth.User{}, err
	}

	return auth.User{
		ID:          int(created.ID),
		FirstName:   created.FirstName,
		LastName:    created.LastName,
		Email:       created.Email,
		Phone:       created.Phone,
		UserType:    created.UserType,
		CreatedAt:   created.CreatedAt.Time,
		UpdatedAt:   created.UpdatedAt.Time,
		LastLoginAt: &created.LastLoginAt.Time,
		IsActive:    created.IsActive.Bool,
	}, nil
}
func (r *Repository) GetUserByEmail(email string) (auth.User, error) {
	user, err := r.db.GetUserByEmail(context.Background(), email)
	if err == sql.ErrNoRows {
		return auth.User{}, sql.ErrNoRows
	}
	if err != nil {
		return auth.User{}, err
	}

	return auth.User{
		ID:           int(user.ID),
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Phone:        user.Phone,
		UserType:     user.UserType,
		CreatedAt:    user.CreatedAt.Time,
		UpdatedAt:    user.UpdatedAt.Time,
		LastLoginAt:  &user.LastLoginAt.Time,
		IsActive:     user.IsActive.Bool,
	}, nil
}

func (r *Repository) UpdateLastLoginAt(userID int) error {
	return r.db.UpdateLastLoginAt(context.Background(), int32(userID))
}

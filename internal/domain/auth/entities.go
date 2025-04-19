package auth

import "time"

type User struct {
	ID           int
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	Phone        string
	UserType     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LastLoginAt  *time.Time
	IsActive     bool
}

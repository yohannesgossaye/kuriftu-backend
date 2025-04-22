package auth

type Repository interface {
	CreateUser(user User) (User, error)
	GetUserByEmail(email string) (User, error)
	UpdateLastLoginAt(userID int) error
}

package auth

type Repository interface {
	CreateUser(user User) (User, error)
}

package auth

type Service interface {
	Register(firstName, lastName, email, password, phone, userType string) (User, error)
}
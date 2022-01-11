package services

type AuthService interface {
	CreateUser()
	VerifyCredential()
	FindByEmail()
}
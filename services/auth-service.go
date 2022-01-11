package services

import (
	"context"
	"rest_api/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthService interface {
	CreateUser(ctx context.Context, user models.User) error
	// VerifyCredential()
	// FindByEmail()
}

type authService struct {
	db *pgxpool.Pool
}

func NewAuthService(db *pgxpool.Pool) AuthService {
	return &authService{db: db}
}

const addUser = `INSERT INTO users(name, email, password, gender_id, create_at, update_at) 
				VALUES($1, $2, $3, $4, now(), now())`

func (a *authService) CreateUser(ctx context.Context, user models.User) error {
	_, err := a.db.Exec(ctx, addUser, user.Name, user.Email, user.Password, user.GenderID)
	if err != nil {
		return err
	}
	return nil
}

package services

import (
	"context"
	"fmt"
	"rest_api/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthService interface {
	CreateUser(ctx context.Context, user models.Register) error
	VerifyCredential(ctx context.Context, user models.Login) error
	FindByEmail(ctx context.Context, email string) string
}

type authService struct {
	db *pgxpool.Pool
}

func NewAuthService(db *pgxpool.Pool) AuthService {
	return &authService{db: db}
}

const addUser = `INSERT INTO users(name, email, password, gender_id, create_at, update_at) VALUES($1, $2, $3, $4, now(), now())`

func (a *authService) CreateUser(ctx context.Context, user models.Register) error {
	duplicate := a.FindByEmail(ctx, user.Email)
	if duplicate == "duplicate" {
		return fmt.Errorf("duplicate")
	}

	user.Password = hasAndSalt([]byte(user.Password))
	_, err := a.db.Exec(ctx, addUser, user.Name, user.Email, user.Password, user.GenderID)
	if err != nil {
		return err
	}
	
	return fmt.Errorf(duplicate)
}

const getByEmail = `SELECT email, password FROM users WHERE email = $1`

func (a *authService) VerifyCredential(ctx context.Context, user models.Login) error {
	pgx := a.db.QueryRow(ctx, getByEmail, user.Email)
	u := new(models.Login)
	err := pgx.Scan(&u.Email, &u.Password)
	if err != nil {
		return err
	}
	compare := comparePwd(u.Password, []byte(user.Password))
	if !compare {
		return fmt.Errorf("invalid credential")
	}
	
	return nil
}

const findByEmail = `SELECT email FROM users WHERE email = $1`

func (a *authService) FindByEmail(ctx context.Context, email string) string {
	pgx := a.db.QueryRow(ctx, findByEmail, email)
	var user models.User
	pgx.Scan(&user.Email)
	if user.Email == email {
		return "duplicate"
	}
	return ""
}


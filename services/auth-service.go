package services

import (
	"context"
	"fmt"
	"rest_api/models"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
)

type AuthService interface {
	CreateUser(ctx context.Context, user models.Register) error
	VerifyCredential(ctx context.Context, user models.Login) (string, error)
}

type authService struct {
	db *pgxpool.Pool
	jwtS JWTService
}

func NewAuthService(db *pgxpool.Pool, jwtS JWTService) AuthService {
	return &authService{db: db,jwtS: jwtS}
}

const addUser = `INSERT INTO users(name, email, password, gender_id, create_at, update_at) VALUES($1, $2, $3, $4, now(), now())`
const userByEmail = `SELECT id FROM users WHERE email = $1`
const updateToken = `UPDATE users SET token = $2 WHERE id = $1`

func (a *authService) CreateUser(ctx context.Context, user models.Register) error {
	duplicate := findByEmail(ctx, a.db, user.Email)
	if duplicate == "duplicate" {
		return fmt.Errorf("duplicate")
	}

	user.Password = hasAndSalt([]byte(user.Password))
	_, err := a.db.Exec(ctx, addUser, user.Name, user.Email, user.Password, user.GenderID)
	if err != nil {
		return err
	}

	var byEmail models.User

	pgx := a.db.QueryRow(ctx, userByEmail, user.Email)
	err = pgx.Scan(&byEmail.ID)
	if err != nil {
		return err
	}

	generateToken := a.jwtS.GenerateToken(strconv.FormatUint(uint64(byEmail.ID), 10))
	_, err = a.db.Exec(ctx, updateToken, byEmail.ID, generateToken)
	if err != nil {
		return err
	}
	
	return fmt.Errorf(duplicate)
}

const getByEmail = `SELECT id, email, password FROM users WHERE email = $1`

func (a *authService) VerifyCredential(ctx context.Context, user models.Login) (string, error) {
	var u models.Login
	
	pgx := a.db.QueryRow(ctx, getByEmail, user.Email)
	err := pgx.Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		return "", err
	}
	
	compare := comparePwd(u.Password, []byte(user.Password))
	if !compare {
		return "", fmt.Errorf("invalid credential")
	}

	// user.Password = hasAndSalt([]byte(user.Password))
	generateToken := a.jwtS.GenerateToken(strconv.FormatUint(uint64(u.ID), 10))
	
	return generateToken, nil
}
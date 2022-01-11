package services

import (
	"context"
	"fmt"
	"rest_api/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserService interface {
	GetById(ctx context.Context, id string) (models.User, error)
	GetAll(ctx context.Context) ([]*models.User, error)
}

type userService struct {
	db *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) UserService {
	return &userService{db: db}
}

const getAll = `SELECT * FROM users`

func (u *userService) GetAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	pgx, err := u.db.Query(ctx, getAll)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	for pgx.Next() {
		var u models.User
		err2 := pgx.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.GenderID, &u.CreateAt, &u.UpdateAt)
		if err2 != nil {
			return nil, fmt.Errorf(err2.Error())
		}
		users = append(users, &u)
	}
	return users, nil
}

const getById = `SELECT * FROM users WHERE id = $1`

func (u *userService) GetById(ctx context.Context, id string) (models.User, error) {
	pgx := u.db.QueryRow(ctx, getById, id)
	var user models.User
	err := pgx.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.GenderID, &user.CreateAt, &user.UpdateAt)
	return user, err
}

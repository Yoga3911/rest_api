package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserService interface {
	GetById()
	GetAll(ctx context.Context) (pgx.Rows, error)
}

type userService struct {
	db *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) UserService {
	return &userService{db: db}
}

func (u *userService) GetById() {

}

const getAll = `SELECT * FROM users`

func (u *userService) GetAll(ctx context.Context) (pgx.Rows, error) {
	pgx, err := u.db.Query(ctx, getAll)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return pgx, nil
}

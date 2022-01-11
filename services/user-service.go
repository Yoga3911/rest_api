package services

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserService interface {
	GetById(ctx context.Context, id string) (pgx.Row)
	GetAll(ctx context.Context) (pgx.Rows, error)
}

type userService struct {
	db *pgxpool.Pool
}

func NewUserService(db *pgxpool.Pool) UserService {
	return &userService{db: db}
}

const getAll = `SELECT * FROM users`

func (u *userService) GetAll(ctx context.Context) (pgx.Rows, error) {
	pgx, err := u.db.Query(ctx, getAll)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return pgx, nil
}

const getById = `SELECT * FROM users WHERE id = $1`

func (u *userService) GetById(ctx context.Context, id string) (pgx.Row) {
	pgx := u.db.QueryRow(ctx, getById, id)
	return pgx
}

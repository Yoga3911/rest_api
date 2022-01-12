package services

import (
	"context"
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

const getAll = `SELECT * FROM users ORDER BY id`

func (u *userService) GetAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	
	pgx, err := u.db.Query(ctx, getAll)
	if err != nil {
		return nil, err
	}

	for pgx.Next() {
		var u models.User
		err = pgx.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.GenderID, &u.CreateAt, &u.UpdateAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}

const getById = `SELECT * FROM users WHERE id = $1`

func (u *userService) GetById(ctx context.Context, id string) (models.User, error) {
	var user models.User
	
	pgx := u.db.QueryRow(ctx, getById, id)
	err := pgx.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.GenderID, &user.CreateAt, &user.UpdateAt)
	
	return user, err
}

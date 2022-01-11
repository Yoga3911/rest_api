package services

import (
	"context"
	"rest_api/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ProfileService interface {
	Update(ctx context.Context, user models.User, id string) error
	Delete(ctx context.Context, id string) error
}

type profileService struct {
	db *pgxpool.Pool
}

func NewProfileService(db *pgxpool.Pool) ProfileService {
	return &profileService{db: db}
}

const updateUser = `UPDATE users SET name = $2, email = $3, password = $4, gender_id = $5, update_at = now() WHERE id = $1`

func (p *profileService) Update(ctx context.Context, user models.User, id string) error {
	_, err := p.db.Exec(ctx, updateUser, id, user.Name, user.Email, user.Password, user.GenderID)
	if err != nil {
		return err
	}
	return nil
}

const deleteUser = `DELETE FROM users WHERE id = $1`

func (p *profileService) Delete(ctx context.Context, id string) error {
	_, err := p.db.Exec(ctx, deleteUser, id)
	if err != nil {
		return err
	}
	return nil
}

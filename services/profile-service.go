package services

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"rest_api/models"
	"strings"
)

type ProfileService interface {
	Update(ctx context.Context, user models.UpdateUser) error
	UpdateByToken(ctx *fiber.Ctx, user models.UpdateUser) error
	Delete(ctx context.Context, id int64) error
	DeleteByToken(ctx *fiber.Ctx) error
}

type profileService struct {
	db   *pgxpool.Pool
	jwtS JWTService
}

func NewProfileService(db *pgxpool.Pool, jwtS JWTService) ProfileService {
	return &profileService{db: db, jwtS: jwtS}
}

const updateUser = `UPDATE users SET name = $2, email = $3, password = $4, gender_id = $5, update_at = now() WHERE id = $1`

func (p *profileService) Update(ctx context.Context, user models.UpdateUser) error {
	user.Password = hasAndSalt([]byte(user.Password))

	_, err := p.db.Exec(ctx, updateUser, user.ID, user.Name, user.Email, user.Password, user.GenderID)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return fmt.Errorf("duplicate")
		}
		return err
	}

	return nil
}

func (p *profileService) UpdateByToken(c *fiber.Ctx, user models.UpdateUser) error {
	autHeader := c.Get("Authorization")
	token, errToken := p.jwtS.ValidateToken(autHeader)
	if errToken != nil {
		log.Println(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	idJWT := claims["id"]

	user.Password = hasAndSalt([]byte(user.Password))

	_, err := p.db.Exec(c.Context(), updateUser, idJWT, user.Name, user.Email, user.Password, user.GenderID)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return fmt.Errorf("duplicate")
		}
		return err
	}

	return nil
}

const deleteUser = `DELETE FROM users WHERE id = $1`

func (p *profileService) Delete(ctx context.Context, id int64) error {
	_, err := p.db.Exec(ctx, deleteUser, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *profileService) DeleteByToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	token, err := p.jwtS.ValidateToken(authHeader)
	if err != nil {
		log.Println(err.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	idJWT := claims["id"]

	_, err = p.db.Exec(c.Context(), deleteUser, idJWT)
	if err != nil {
		log.Println(err.Error())
	}

	return nil
}

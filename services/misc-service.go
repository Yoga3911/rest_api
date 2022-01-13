package services

import (
	"context"
	"fmt"
	"rest_api/models"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type Validator struct {
	Failed string
	Tag string
	Value interface{}
}

func StructValidator(user interface{}) []*Validator {
	var errors []*Validator
	validated := validator.New()
	err := validated.Struct(user)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			var er Validator
			er.Failed = e.StructNamespace()
			er.Tag = e.Tag()
			er.Value = e.Param()
			errors = append(errors, &er)
		}
	}
	return errors
}

const find = `SELECT email FROM users WHERE email = $1`

func findByEmail(ctx context.Context, db *pgxpool.Pool, email string) string {
	var user models.User
	
	pgx := db.QueryRow(ctx, find, email)
	pgx.Scan(&user.Email)
	if user.Email == email {
		return "duplicate"
	}

	return ""
}

func hasAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		panic(err.Error())
	}
	
	return string(hash)
}

func comparePwd(hashPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
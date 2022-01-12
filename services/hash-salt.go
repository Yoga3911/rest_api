package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

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
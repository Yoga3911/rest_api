package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userId string) string
	ValidateToken(token string) (*jwt.Token, error)
	// GetTokenValues(token string)
}

type jwtCustomClaim struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "qwerty",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "qwerty"
	}

	return secretKey
}

func (j *jwtService) GenerateToken(ID string) string {
	claims := &jwtCustomClaim{
		ID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Println(err.Error())
	}

	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})

}

// func (j *jwtService) GetTokenValues(token string) {
// 	claims := jwt.MapClaims{}
// 	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
// 		return []byte(j.secretKey), nil
// 	})
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
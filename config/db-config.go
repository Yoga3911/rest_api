package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	// "github.com/joho/godotenv"
	"os"
)

func DatabaseConnection() *pgxpool.Pool {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err)
	// }

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",dbHost, dbUser, dbPass, dbPort, dbName)

	pgx, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}

	return pgx
}

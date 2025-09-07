package services

import (
	"os"
	"plataforma-cursos/pkg/database"

	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load(".env")
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		panic("DATABASE_URL não definida para testes")
	}
	database.Connect(dbURL)
}

package main

import (
	"log"
	"net/http"
	"os"
	"plataforma-cursos/internal/routes"
	"plataforma-cursos/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL não definida")
	}
	database.Connect(dbURL)
	defer database.Close()

	router := gin.New()
	routes.SetupRoutes(router)
	log.Println("Iniciando o servidor na porta 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

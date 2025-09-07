package routes

import (
	"plataforma-cursos/internal/di"
	"plataforma-cursos/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(middleware.ExceptionHandler())

	courseController := di.InitializeCourseController()
	userController := di.InitializeUserController()

	// Rotas públicas
	router.POST("/users", userController.CreateUser)
	router.GET("/courses/:id", courseController.GetCourse)

	// Rotas protegidas (exigem autenticação)
	auth := router.Group("/")
	auth.Use(middleware.AuthMiddleware())
	auth.POST("/courses", courseController.CreateCourse)
	auth.PUT("/courses/:id", courseController.UpdateCourse)
	auth.DELETE("/courses/:id", courseController.DeleteCourse)
	auth.GET("/users/:id", userController.GetUser)
	auth.PUT("/users/:id", userController.UpdateUser)
	auth.DELETE("/users/:id", userController.DeleteUser)
}

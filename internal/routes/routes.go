package routes

import (
    "github.com/gin-gonic/gin"
    "plataforma-cursos/internal/controllers"
    "plataforma-cursos/internal/services"
    "plataforma-cursos/internal/middleware"
)

func SetupRoutes(router *gin.Engine) {
    courseService := services.NewCourseService()
    courseController := controllers.NewCourseController(*courseService)

    userService := services.NewUserService()
    userController := controllers.NewUserController(*userService)

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
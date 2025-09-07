//go:build wireinject
// +build wireinject

package di

import (
	"database/sql"
	"plataforma-cursos/internal/controllers"
	"plataforma-cursos/internal/services"
	"plataforma-cursos/pkg/database"

	"github.com/google/wire"
)

func ProvideDB() *sql.DB {
	return database.GetDB()
}

func InitializeCourseController() *controllers.CourseController {
	wire.Build(
		ProvideDB,
		services.NewCourseService,
		controllers.NewCourseController,
	)
	return nil
}

func InitializeUserController() *controllers.UserController {
	wire.Build(
		ProvideDB,
		services.NewUserService,
		controllers.NewUserController,
	)
	return nil
}

package controllers

import (
	"fmt"
	"net/http"
	"plataforma-cursos/internal/models"
	"plataforma-cursos/internal/services"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	Service *services.CourseService
}

func NewCourseController(service *services.CourseService) *CourseController {
	return &CourseController{Service: service}
}

func (c *CourseController) CreateCourse(ctx *gin.Context) {
	var course models.Course
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Service.AddCourse(course)
	ctx.JSON(http.StatusCreated, course)
}

func (c *CourseController) GetCourse(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id int
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	course, err := c.Service.FindCourse(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	ctx.JSON(http.StatusOK, course)
}

func (c *CourseController) UpdateCourse(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id int
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var course models.Course
	if err := ctx.ShouldBindJSON(&course); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	course.ID = id
	if err := c.Service.ModifyCourse(course); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	ctx.JSON(http.StatusOK, course)
}

func (c *CourseController) DeleteCourse(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id int
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.Service.RemoveCourse(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

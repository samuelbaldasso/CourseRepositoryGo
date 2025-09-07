package controllers

import (
	"net/http"
	"fmt"
	"github.com/gin-gonic/gin"
	"plataforma-cursos/internal/models"
	"plataforma-cursos/internal/services"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := c.service.AddUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id int
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	user, err := c.service.FindUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id int
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	if err := c.service.ModifyUser(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id int
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.service.RemoveUser(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

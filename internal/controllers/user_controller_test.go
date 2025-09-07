package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"plataforma-cursos/internal/models"
	"plataforma-cursos/internal/services"
	"github.com/gin-gonic/gin"
	"fmt"
)

func setupUserController() (*UserController, *services.UserService) {
	svc := services.NewUserService()
	ctrl := NewUserController(*svc)
	return ctrl, svc
}

func TestCreateUserController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl, _ := setupUserController()
	r := gin.New()
	r.POST("/users", ctrl.CreateUser)

	user := models.User{Name: "Teste", Email: "teste@email.com"}
	body, _ := json.Marshal(user)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("esperava status 201, obteve %d", w.Code)
	}
}

func TestGetUserController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl, svc := setupUserController()
	r := gin.New()
	r.GET("/users/:id", ctrl.GetUser)
	created, _ := svc.AddUser(models.User{Name: "Teste", Email: "teste@email.com"})
	w := httptest.NewRecorder()
	url := "/users/" + fmt.Sprint(created.ID)
	req, _ := http.NewRequest("GET", url, nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("esperava status 200, obteve %d", w.Code)
	}
}

func TestUpdateUserController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl, svc := setupUserController()
	r := gin.New()
	r.PUT("/users/:id", ctrl.UpdateUser)
	created, _ := svc.AddUser(models.User{Name: "Teste", Email: "teste@email.com"})
	created.Name = "Novo Nome"
	body, _ := json.Marshal(created)
	w := httptest.NewRecorder()
	url := "/users/" + fmt.Sprint(created.ID)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("esperava status 200, obteve %d", w.Code)
	}
}

func TestDeleteUserController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl, svc := setupUserController()
	r := gin.New()
	r.DELETE("/users/:id", ctrl.DeleteUser)
	created, _ := svc.AddUser(models.User{Name: "Teste", Email: "teste@email.com"})
	w := httptest.NewRecorder()
	url := "/users/" + fmt.Sprint(created.ID)
	req, _ := http.NewRequest("DELETE", url, nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("esperava status 204, obteve %d", w.Code)
	}
}

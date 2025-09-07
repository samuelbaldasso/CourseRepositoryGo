package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"plataforma-cursos/internal/controllers"
	"plataforma-cursos/internal/di"
	"plataforma-cursos/internal/models"
	"strconv"
	"testing"

	_ "plataforma-cursos/internal/services"

	"github.com/gin-gonic/gin"
)

func setupUserController() *controllers.UserController {
	ctrl := di.InitializeUserController()
	if ctrl == nil || ctrl.Service == nil {
		panic("UserController ou Service está nil. Verifique a configuração do DI e a conexão com o banco.")
	}
	return ctrl
}

func TestCreateUserController(t *testing.T) {
	ctrl := setupUserController()
	gin.SetMode(gin.TestMode)
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
	ctrl := setupUserController()
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/users/:id", ctrl.GetUser)
	created, _ := ctrl.Service.AddUser(models.User{Name: "Teste", Email: "teste@email.com"})
	w := httptest.NewRecorder()
	url := "/users/" + strconv.Itoa(created.ID)
	req, _ := http.NewRequest("GET", url, nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("esperava status 200, obteve %d", w.Code)
	}
}

func TestUpdateUserController(t *testing.T) {
	ctrl := setupUserController()
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.PUT("/users/:id", ctrl.UpdateUser)
	created, _ := ctrl.Service.AddUser(models.User{Name: "Teste", Email: "teste@email.com"})
	created.Name = "Novo Nome"
	body, _ := json.Marshal(created)
	w := httptest.NewRecorder()
	url := "/users/" + strconv.Itoa(created.ID)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("esperava status 200, obteve %d", w.Code)
	}
}

func TestDeleteUserController(t *testing.T) {
	ctrl := setupUserController()
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.DELETE("/users/:id", ctrl.DeleteUser)
	created, _ := ctrl.Service.AddUser(models.User{Name: "Teste", Email: "teste@email.com"})
	w := httptest.NewRecorder()
	url := "/users/" + strconv.Itoa(created.ID)
	req, _ := http.NewRequest("DELETE", url, nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("esperava status 204, obteve %d", w.Code)
	}
}

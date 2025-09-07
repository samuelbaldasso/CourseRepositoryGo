package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"plataforma-cursos/internal/di"
	"plataforma-cursos/internal/models"
	"strconv"
	"testing"

	_ "plataforma-cursos/internal/services"

	"github.com/gin-gonic/gin"
)

func TestCreateCourseController(t *testing.T) {
	ctrl := di.InitializeCourseController()
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/courses", ctrl.CreateCourse)

	course := models.Course{Title: "Go", Description: "Curso Go", Duration: 10}
	body, _ := json.Marshal(course)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/courses", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("esperava status 201, obteve %d", w.Code)
	}
}

func TestGetCourseController(t *testing.T) {
	ctrl := di.InitializeCourseController()
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/courses/:id", ctrl.GetCourse)
	created, _ := ctrl.Service.AddCourse(models.Course{Title: "Go", Description: "Curso Go", Duration: 10})
	w := httptest.NewRecorder()
	url := "/courses/" + strconv.Itoa(created.ID)
	req, _ := http.NewRequest("GET", url, nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("esperava status 200, obteve %d", w.Code)
	}
}

func TestUpdateCourseController(t *testing.T) {
	ctrl := di.InitializeCourseController()
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.PUT("/courses/:id", ctrl.UpdateCourse)
	created, _ := ctrl.Service.AddCourse(models.Course{Title: "Go", Description: "Curso Go", Duration: 10})
	mod := models.Course{Title: "Go Avançado", Description: "Avançado", Duration: 20}
	body, _ := json.Marshal(mod)
	w := httptest.NewRecorder()
	url := "/courses/" + strconv.Itoa(created.ID)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("esperava status 200, obteve %d", w.Code)
	}
}

func TestDeleteCourseController(t *testing.T) {
	ctrl := di.InitializeCourseController()
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.DELETE("/courses/:id", ctrl.DeleteCourse)
	created, _ := ctrl.Service.AddCourse(models.Course{Title: "Go", Description: "Curso Go", Duration: 10})
	w := httptest.NewRecorder()
	url := "/courses/" + strconv.Itoa(created.ID)
	req, _ := http.NewRequest("DELETE", url, nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("esperava status 204, obteve %d", w.Code)
	}
}

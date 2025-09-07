package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"plataforma-cursos/internal/models"
	"plataforma-cursos/internal/services"
	"github.com/gin-gonic/gin"
)

func setupCourseController() (*CourseController, *services.CourseService) {
	svc := services.NewCourseService()
	ctrl := NewCourseController(*svc)
	return ctrl, svc
}

func TestCreateCourseController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl, _ := setupCourseController()
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
	gin.SetMode(gin.TestMode)
	ctrl, svc := setupCourseController()
	r := gin.New()
	r.GET("/courses/:id", ctrl.GetCourse)
	_, _ = svc.AddCourse(models.Course{Title: "Go", Description: "Curso Go", Duration: 10})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/courses/1", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("esperava status 200, obteve %d", w.Code)
	}
}

func TestUpdateCourseController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl, svc := setupCourseController()
	r := gin.New()
	r.PUT("/courses/:id", ctrl.UpdateCourse)
	created, _ := svc.AddCourse(models.Course{Title: "Go", Description: "Curso Go", Duration: 10})
	mod := models.Course{Title: "Go Avançado", Description: "Avançado", Duration: 20}
	body, _ := json.Marshal(mod)
	w := httptest.NewRecorder()
	url := "/courses/" + fmt.Sprint(created.ID)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("esperava status 200, obteve %d", w.Code)
	}
}

func TestDeleteCourseController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl, svc := setupCourseController()
	r := gin.New()
	r.DELETE("/courses/:id", ctrl.DeleteCourse)
	created, _ := svc.AddCourse(models.Course{Title: "Go", Description: "Curso Go", Duration: 10})
	w := httptest.NewRecorder()
	url := "/courses/" + fmt.Sprint(created.ID)
	req, _ := http.NewRequest("DELETE", url, nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("esperava status 204, obteve %d", w.Code)
	}
}

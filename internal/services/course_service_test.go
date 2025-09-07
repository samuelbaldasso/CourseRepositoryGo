package services

import (
	"testing"
	"plataforma-cursos/internal/models"
)

func TestAddCourse(t *testing.T) {
	svc := NewCourseService()
	course := models.Course{Title: "Go", Description: "Curso Go", Duration: 10}
	created, err := svc.AddCourse(course)
	if err != nil {
		t.Fatalf("erro ao adicionar curso: %v", err)
	}
	if created.ID == 0 {
		t.Error("ID não foi atribuído")
	}
}

func TestFindCourse(t *testing.T) {
	svc := NewCourseService()
	course := models.Course{Title: "Go", Description: "Curso Go", Duration: 10}
	created, _ := svc.AddCourse(course)
	found, err := svc.FindCourse(created.ID)
	if err != nil {
		t.Fatalf("erro ao buscar curso: %v", err)
	}
	if found.Title != course.Title {
		t.Error("curso retornado incorreto")
	}
}

func TestModifyCourse(t *testing.T) {
	svc := NewCourseService()
	course := models.Course{Title: "Go", Description: "Curso Go", Duration: 10}
	created, _ := svc.AddCourse(course)
	mod := models.Course{ID: created.ID, Title: "Go Avançado", Description: "Avançado", Duration: 20}
	err := svc.ModifyCourse(mod)
	if err != nil {
		t.Fatalf("erro ao modificar curso: %v", err)
	}
	found, _ := svc.FindCourse(created.ID)
	if found.Title != "Go Avançado" {
		t.Error("modificação não persistiu")
	}
}

func TestRemoveCourse(t *testing.T) {
	svc := NewCourseService()
	course := models.Course{Title: "Go", Description: "Curso Go", Duration: 10}
	created, _ := svc.AddCourse(course)
	err := svc.RemoveCourse(created.ID)
	if err != nil {
		t.Fatalf("erro ao remover curso: %v", err)
	}
	_, err = svc.FindCourse(created.ID)
	if err == nil {
		t.Error("curso não foi removido")
	}
}

package services_test

import (
	"plataforma-cursos/internal/di"
	"plataforma-cursos/internal/models"
	_ "plataforma-cursos/internal/services"
	"testing"
)

func TestAddUser(t *testing.T) {
	svc := di.InitializeUserController().Service
	user := models.User{Name: "João", Email: "joao@email.com"}
	created, err := svc.AddUser(user)
	if err != nil {
		t.Fatalf("erro ao adicionar usuário: %v", err)
	}
	if created.ID == 0 {
		t.Error("ID não foi atribuído")
	}
}

func TestAddUser_DuplicateEmail(t *testing.T) {
	svc := di.InitializeUserController().Service
	user := models.User{Name: "João", Email: "joao@email.com"}
	_, _ = svc.AddUser(user)
	_, err := svc.AddUser(user)
	if err == nil {
		t.Error("esperava erro de usuário duplicado")
	}
}

func TestFindUser(t *testing.T) {
	svc := di.InitializeUserController().Service
	user := models.User{Name: "Maria", Email: "maria@email.com"}
	created, _ := svc.AddUser(user)
	found, err := svc.FindUser(created.ID)
	if err != nil {
		t.Fatalf("erro ao buscar usuário: %v", err)
	}
	if found.Email != user.Email {
		t.Error("usuário retornado incorreto")
	}
}

func TestModifyUser(t *testing.T) {
	svc := di.InitializeUserController().Service
	user := models.User{Name: "Ana", Email: "ana@email.com"}
	created, _ := svc.AddUser(user)
	created.Name = "Ana Paula"
	err := svc.ModifyUser(created)
	if err != nil {
		t.Fatalf("erro ao modificar usuário: %v", err)
	}
	found, _ := svc.FindUser(created.ID)
	if found.Name != "Ana Paula" {
		t.Error("modificação não persistiu")
	}
}

func TestRemoveUser(t *testing.T) {
	svc := di.InitializeUserController().Service
	user := models.User{Name: "Carlos", Email: "carlos@email.com"}
	created, _ := svc.AddUser(user)
	err := svc.RemoveUser(created.ID)
	if err != nil {
		t.Fatalf("erro ao remover usuário: %v", err)
	}
	_, err = svc.FindUser(created.ID)
	if err == nil {
		t.Error("usuário não foi removido")
	}
}

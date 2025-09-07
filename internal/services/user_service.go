package services

import (
	"errors"
	"sync"
	"plataforma-cursos/internal/models"
)

type UserService struct {
	users  map[int]models.User
	mu     sync.Mutex
	nextID int
}

func NewUserService() *UserService {
	return &UserService{
		users:  make(map[int]models.User),
		nextID: 1,
	}
}

func (s *UserService) AddUser(user models.User) (models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user.ID = s.nextID
	s.nextID++

	// Email único
	for _, u := range s.users {
		if u.Email == user.Email {
			return models.User{}, errors.New("user already exists")
		}
	}

	s.users[user.ID] = user
	return user, nil
}

func (s *UserService) FindUser(id int) (models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) ModifyUser(user models.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[user.ID]; !exists {
		return errors.New("user not found")
	}
	// Email único
	for _, u := range s.users {
		if u.Email == user.Email && u.ID != user.ID {
			return errors.New("email already in use")
		}
	}

	s.users[user.ID] = user
	return nil
}

func (s *UserService) RemoveUser(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(s.users, id)
	return nil
}

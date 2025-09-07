package services

import (
	"errors"
	"sync"
	"plataforma-cursos/internal/models"
)

type CourseService struct {
	courses map[int]models.Course
	mu      sync.Mutex
	nextID  int
}

func NewCourseService() *CourseService {
	return &CourseService{
		courses: make(map[int]models.Course),
		nextID:  1,
	}
}

func (s *CourseService) AddCourse(course models.Course) (models.Course, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	course.ID = s.nextID
	s.nextID++

	if _, exists := s.courses[course.ID]; exists {
		return models.Course{}, errors.New("course already exists")
	}
	s.courses[course.ID] = course
	return course, nil
}

func (s *CourseService) FindCourse(id int) (models.Course, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	course, exists := s.courses[id]
	if !exists {
		return models.Course{}, errors.New("course not found")
	}
	return course, nil
}

func (s *CourseService) ModifyCourse(course models.Course) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.courses[course.ID]; !exists {
		return errors.New("course not found")
	}
	s.courses[course.ID] = course
	return nil
}

func (s *CourseService) RemoveCourse(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.courses[id]; !exists {
		return errors.New("course not found")
	}
	delete(s.courses, id)
	return nil
}
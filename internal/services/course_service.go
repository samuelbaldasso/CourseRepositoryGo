package services

import (
	"database/sql"
	"errors"
	"plataforma-cursos/internal/models"
)

type CourseService struct {
	db *sql.DB
}

func NewCourseService(db *sql.DB) *CourseService {
	return &CourseService{db: db}
}

func (s *CourseService) AddCourse(course models.Course) (models.Course, error) {
	query := `INSERT INTO courses (title, description, duration) VALUES ($1, $2, $3) RETURNING id`
	err := s.db.QueryRow(query, course.Title, course.Description, course.Duration).Scan(&course.ID)
	if err != nil {
		return models.Course{}, err
	}
	return course, nil
}

func (s *CourseService) FindCourse(id int) (models.Course, error) {
	var course models.Course
	query := `SELECT id, title, description, duration FROM courses WHERE id = $1`
	row := s.db.QueryRow(query, id)
	err := row.Scan(&course.ID, &course.Title, &course.Description, &course.Duration)
	if err == sql.ErrNoRows {
		return models.Course{}, errors.New("course not found")
	} else if err != nil {
		return models.Course{}, err
	}
	return course, nil
}

func (s *CourseService) ModifyCourse(course models.Course) error {
	query := `UPDATE courses SET title=$1, description=$2, duration=$3 WHERE id=$4`
	res, err := s.db.Exec(query, course.Title, course.Description, course.Duration, course.ID)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("course not found")
	}
	return nil
}

func (s *CourseService) RemoveCourse(id int) error {
	query := `DELETE FROM courses WHERE id=$1`
	res, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("course not found")
	}
	return nil
}

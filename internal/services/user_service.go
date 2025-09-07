package services

import (
	"database/sql"
	"errors"
	"plataforma-cursos/internal/models"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) AddUser(user models.User) (models.User, error) {
	// Email único
	var exists int
	err := s.db.QueryRow("SELECT COUNT(1) FROM users WHERE email = $1", user.Email).Scan(&exists)
	if err != nil {
		return models.User{}, err
	}
	if exists > 0 {
		return models.User{}, errors.New("user already exists")
	}
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err = s.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) FindUser(id int) (models.User, error) {
	var user models.User
	query := `SELECT id, name, email FROM users WHERE id = $1`
	row := s.db.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		return models.User{}, errors.New("user not found")
	} else if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (s *UserService) ModifyUser(user models.User) error {
	// Email único
	var exists int
	err := s.db.QueryRow("SELECT COUNT(1) FROM users WHERE email = $1 AND id != $2", user.Email, user.ID).Scan(&exists)
	if err != nil {
		return err
	}
	if exists > 0 {
		return errors.New("email already in use")
	}
	query := `UPDATE users SET name=$1, email=$2 WHERE id=$3`
	res, err := s.db.Exec(query, user.Name, user.Email, user.ID)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (s *UserService) RemoveUser(id int) error {
	query := `DELETE FROM users WHERE id=$1`
	res, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("user not found")
	}
	return nil
}

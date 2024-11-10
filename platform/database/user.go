package database

import (
	"database/sql"
	"fmt"
	"web-lab2/platform/models"

	_ "github.com/lib/pq"
)

func (s *PostgresStorage) CreateUser(user *models.UserDto) (string, error) {
	query := "INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id"
	id := ""
	s.db.QueryRow(query, user.Username, user.Password).Scan(&id)
	return id, nil
}

func (s *PostgresStorage) GetUser(user *models.UserDto) (*models.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = $1 AND password = $2"
	u := new(models.User)
	err := s.db.QueryRow(query, user.Username, user.Password).Scan(&u.ID, &u.Username, &u.Password)
	return u, err
}

func (s *PostgresStorage) GetUserByID(id string) (*models.User, error) {
	query := "SELECT id, username, password FROM users WHERE id = $1"
	u := new(models.User)
	err := s.db.QueryRow(query, id).Scan(&u.ID, &u.Username, &u.Password)
	return u, err
}

func (s *PostgresStorage) DeleteUser(id string) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := s.db.Exec(query, id)
	return err
}

func (s *PostgresStorage) UsernameExists(username string) bool {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)"
	exists := false
	s.db.QueryRow(query, username).Scan(&exists)
	return exists
}

func (s *PostgresStorage) SimGetUser(user *models.SqlInjUser) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s' AND password = '%s'", user.Username, user.Password)
	rows, err := s.db.Query(query)
	return rows, err
}

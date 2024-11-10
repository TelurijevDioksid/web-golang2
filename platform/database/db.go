package database

import (
	"database/sql"
	"web-lab2/platform/models"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateUser(user *models.UserDto) error
	GetUser(user *models.UserDto) (*models.User, error)
	DeleteUser(id string) error
	UsernameExists(username string) bool
	SimGetUser(user *models.SqlInjUser) (*sql.Rows, error)
}

type PostgresStorage struct {
	db *sql.DB
}

func New(conStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", conStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Setup() error {
	query := `
        CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

        CREATE TABLE IF NOT EXISTS users (
            id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
            username VARCHAR(100) NOT NULL,
            password VARCHAR(100) NOT NULL
        );

        DELETE FROM users WHERE true;

        INSERT INTO users (username, password) VALUES ('ime1', 'lozinka1');
        INSERT INTO users (username, password) VALUES ('ime2', 'lozinka2');
        INSERT INTO users (username, password) VALUES ('ime3', 'lozinka3');
        INSERT INTO users (username, password) VALUES ('ime4', 'lozinka4');
    `
	_, err := s.db.Exec(query)
	return err
}

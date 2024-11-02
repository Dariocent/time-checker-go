package user

import (
	"database/sql"
	"errors"

	_ "github.com/lib/pq" // PostgreSQL driver
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserStore interface {
	CreateUser(username, password, email string) error
	GetAllUsers() ([]User, error)
	DeleteUser(username string) error
}

type SQLUserStore struct {
	DB *sql.DB
}

func (s *SQLUserStore) CreateUser(username, password, email string) error {
	//check if the user already exists
	var count int
	err := s.DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1", username).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("user already exists")
	}
	_, err = s.DB.Exec("INSERT INTO users (username, password, email) VALUES ($1, $2, $3)", username, password, email)
	return err
}

func (s *SQLUserStore) DeleteUser(username string) error {
	//check if the user exists
	var count int
	err := s.DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = $1", username).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("user does not exist")
	}
	_, err = s.DB.Exec("DELETE FROM users WHERE username = $1", username)
	return err
}

func (s *SQLUserStore) GetAllUsers() ([]User, error) {
	rows, err := s.DB.Query("SELECT id, username, password, email FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

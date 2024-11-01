package user

import (
	"database/sql"

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
}

type SQLUserStore struct {
	DB *sql.DB
}

func (s *SQLUserStore) CreateUser(username, password, email string) error {
	_, err := s.DB.Exec("INSERT INTO users (username, password, email) VALUES ($1, $2, $3)", username, password, email)
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

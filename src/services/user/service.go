package user

import "errors"

type UserService struct {
	UserStore UserStore
}

func (s *UserService) CreateUser(username, password, email string) error {
	//check payload is correct
	if username == "" || password == "" || email == "" {
		return errors.New("username, password and email are required")
	}
	return s.UserStore.CreateUser(username, password, email)
}

func (s *UserService) DeleteUser(username string) error {
	//check payload is correct
	if username == "" {
		return errors.New("username is required")
	}
	return s.UserStore.DeleteUser(username)
}

func (s *UserService) ListUsers() ([]User, error) {
	return s.UserStore.GetAllUsers()
}

package user

type UserService struct {
	UserStore UserStore
}

func (s *UserService) CreateUser(username, password, email string) error {
	return s.UserStore.CreateUser(username, password, email)
}

func (s *UserService) ListUsers() ([]User, error) {
	return s.UserStore.GetAllUsers()
}

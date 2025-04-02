package core

import "fmt"

type UserService interface {
	RegisterUser(user User) error
}
type userServiceImpl struct {
	repo UserRepository
}

func (o *userServiceImpl) RegisterUser(user User) error {

	if user.Username == "" || user.Password == "" || user.Email == "" {
		return fmt.Errorf("username, password, and email are required")
	}
	if err := o.repo.Save(user); err != nil {
		return err
	}
	return nil
}

func NewOrderService(repo UserRepository) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}

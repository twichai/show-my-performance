package core

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user User) error
}
type userServiceImpl struct {
	repo UserRepository
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (o *userServiceImpl) RegisterUser(user User) error {

	if user.Username == "" || user.Password == "" || user.Email == "" {
		return fmt.Errorf("username, password, and email are required")
	}
	if len(user.Password) < 6 {
		return fmt.Errorf("password must be at least 6 characters long")
	}
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %v", err)
	}
	user.Password = hashedPassword

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

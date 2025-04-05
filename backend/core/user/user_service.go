package userCore

import (
	"fmt"
	"show-my-performance/backend/model"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user model.User) error
	Login(user model.User) (*model.User, error)
}
type userServiceImpl struct {
	repo UserRepository
}

// Login implements UserService.
func (o *userServiceImpl) Login(user model.User) (*model.User, error) {
	if user.Email == "" || user.Password == "" {
		return nil, fmt.Errorf("email and password are required")
	}

	existingUser, err := o.repo.GetByEmail(user.Email)
	if err != nil || existingUser == nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	return existingUser, nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (o *userServiceImpl) RegisterUser(user model.User) error {

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

func NewUserService(repo UserRepository) UserService {
	return &userServiceImpl{
		repo: repo,
	}
}

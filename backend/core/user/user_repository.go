package userCore

import "show-my-performance/backend/model"

type UserRepository interface {
	Save(user model.User) error
	GetByEmail(userEmail string) (*model.User, error)
}

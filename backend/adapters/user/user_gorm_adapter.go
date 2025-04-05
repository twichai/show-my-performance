package userAdapter

import (
	userCore "show-my-performance/backend/core/user"
	"show-my-performance/backend/model"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

// GetByEmail implements core.UserRepository.
func (r *GormUserRepository) GetByEmail(userEmail string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", userEmail).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) userCore.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Save(user model.User) error {
	if result := r.db.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

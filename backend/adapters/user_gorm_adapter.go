package adapters

import (
	"show-my-performance/backend/core"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

// GetByEmail implements core.UserRepository.
func (r *GormUserRepository) GetByEmail(userEmail string) (*core.User, error) {
	var user core.User
	if err := r.db.Where("email = ?", userEmail).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) core.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Save(user core.User) error {
	if result := r.db.Create(&user); result.Error != nil {
		return result.Error
	}
	return nil
}

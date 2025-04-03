package core

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string `json:"username"`
	Password        string `json:"password"`
	ProfileImageURL string `json:"profile_image_url"`
	Email           string `json:"email" gorm:"uniqueIndex"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	PhoneNumber     string `json:"phone_number"`
	Posts           []Post `json:"posts" gorm:"foreignKey:UserID"`
}

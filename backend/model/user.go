package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username        string `json:"username"`
	Password        string `json:"password"`
	ProfileImageURL string `json:"pofileImageUrl"`
	Email           string `json:"email" gorm:"uniqueIndex"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	PhoneNumber     string `json:"phoneNumber"`
	Posts           []Post `json:"posts" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

var CurrentUser User

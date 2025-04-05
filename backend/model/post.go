package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string `json:"title" form:"title"`
	Content  string `json:"content" form:"content"`
	ImageUrl string `json:"imageUrl" form:"ImageUrl"`
	UserID   uint   `json:"UserId" form:"UserId"`
	User     User   `json:"user"`
}

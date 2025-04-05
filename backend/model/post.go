package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string `json:"title" form:"title"`
	Content  string `json:"content" form:"content"`
	ImageUrl string `json:"image_url" form:"image_url"`
	UserID   uint   `json:"user_id" form:"user_id"`
}

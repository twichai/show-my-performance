package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title     string     `json:"title" form:"title"`
	Content   string     `json:"content" form:"content"`
	ImageUrl  string     `json:"imageUrl" form:"ImageUrl"`
	UserID    uint       `json:"UserId" form:"UserId"`
	User      User       `json:"user"`
	Reactions []Reaction `json:"reactions" gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
}

type PostDto struct {
	ID        uint       `json:"id" form:"id"`
	CreatedAt string     `json:"createdAt" form:"createdAt"`
	UpdatedAt string     `json:"updatedAt" form:"updatedAt"`
	DeletedAt string     `json:"deletedAt" form:"deletedAt"`
	Title     string     `json:"title" form:"title"`
	Content   string     `json:"content" form:"content"`
	ImageUrl  string     `json:"imageUrl" form:"ImageUrl"`
	UserID    uint       `json:"UserId" form:"UserId"`
	User      UserDto    `json:"user"`
	Reactions []Reaction `json:"reactions" gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE"`
}

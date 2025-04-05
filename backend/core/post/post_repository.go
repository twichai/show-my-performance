package postCore

import "show-my-performance/backend/model"

type PostRepository interface {
	GetAllPosts() ([]model.Post, error)
	GetPostByID(id uint) (*model.Post, error)
	CreatePost(post *model.Post) error
	UpdatePost(post *model.Post) error
	DeletePost(id uint, userID uint) error
	GetPostsByUserID(userID uint) ([]model.Post, error)
}

package adapters

import (
	"show-my-performance/backend/core"

	"gorm.io/gorm"
)

type postGormRepository struct {
	db *gorm.DB
}

// CreatePost implements core.PostRepository.
func (p *postGormRepository) CreatePost(post *core.Post) error {
	if result := p.db.Create(&post); result.Error != nil {
		return result.Error
	}
	return nil
}

// DeletePost implements core.PostRepository.
func (p *postGormRepository) DeletePost(id uint) error {
	panic("unimplemented")
}

// GetAllPosts implements core.PostRepository.
func (p *postGormRepository) GetAllPosts() ([]core.Post, error) {
	panic("unimplemented")
}

// GetPostByID implements core.PostRepository.
func (p *postGormRepository) GetPostByID(id uint) (*core.Post, error) {
	post := &core.Post{}
	if p.db.First(post, id).Error != nil {
		return nil, p.db.Error
	}
	return post, nil
}

// GetPostsByUserID implements core.PostRepository.
func (p *postGormRepository) GetPostsByUserID(userID uint) ([]core.Post, error) {
	panic("unimplemented")
}

// UpdatePost implements core.PostRepository.
func (p *postGormRepository) UpdatePost(post *core.Post) error {
	panic("unimplemented")
}

func NewGormPostRepository(db *gorm.DB) core.PostRepository {
	return &postGormRepository{db: db}
}

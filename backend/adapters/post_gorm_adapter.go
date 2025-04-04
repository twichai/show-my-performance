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
func (p *postGormRepository) DeletePost(id uint, userID uint) error {
	if result := p.db.Delete(&core.Post{}, id).Where(userID); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllPosts implements core.PostRepository.
func (p *postGormRepository) GetAllPosts() ([]core.Post, error) {
	post := []core.Post{}
	if p.db.Order("updated_at desc").Limit(10).Find(&post).Error != nil {
		return nil, p.db.Error
	}
	return post, nil
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
	posts := []core.Post{}
	if p.db.Where("user_id = ?", userID).Find(&posts).Error != nil {
		return nil, p.db.Error
	}
	return posts, nil
}

// UpdatePost implements core.PostRepository.
func (p *postGormRepository) UpdatePost(post *core.Post) error {
	if result := p.db.Save(post); result.Error != nil {
		return result.Error
	}
	return nil
}

func NewGormPostRepository(db *gorm.DB) core.PostRepository {
	return &postGormRepository{db: db}
}

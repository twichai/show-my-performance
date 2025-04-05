package postAdapter

import (
	postCore "show-my-performance/backend/core/post"
	"show-my-performance/backend/model"

	"gorm.io/gorm"
)

type postGormRepository struct {
	db *gorm.DB
}

// CreatePost implements core.PostRepository.
func (p *postGormRepository) CreatePost(post *model.Post) error {
	if result := p.db.Create(&post); result.Error != nil {
		return result.Error
	}
	return nil
}

// DeletePost implements core.PostRepository.
func (p *postGormRepository) DeletePost(id uint, userID uint) error {
	if result := p.db.Delete(&model.Post{}, id).Where(userID); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllPosts implements core.PostRepository.
func (p *postGormRepository) GetAllPosts() ([]model.Post, error) {
	posts := []model.Post{}
	if err := p.db.Preload("User").Order("updated_at desc").Limit(10).Find(&posts).Error; err != nil {
		return nil, p.db.Error
	}
	return posts, nil
}

// GetPostByID implements core.PostRepository.
func (p *postGormRepository) GetPostByID(id uint) (*model.Post, error) {
	post := &model.Post{}
	if p.db.First(post, id).Error != nil {
		return nil, p.db.Error
	}
	return post, nil
}

// GetPostsByUserID implements core.PostRepository.
func (p *postGormRepository) GetPostsByUserID(userID uint) ([]model.Post, error) {
	posts := []model.Post{}
	if p.db.Where("user_id = ?", userID).Find(&posts).Error != nil {
		return nil, p.db.Error
	}
	return posts, nil
}

// UpdatePost implements core.PostRepository.
func (p *postGormRepository) UpdatePost(post *model.Post) error {
	if result := p.db.Save(post); result.Error != nil {
		return result.Error
	}
	return nil
}

func NewGormPostRepository(db *gorm.DB) postCore.PostRepository {
	return &postGormRepository{db: db}
}

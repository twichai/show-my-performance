package core

import "fmt"

type PostService interface {
	GetAllPosts() ([]Post, error)
	GetPostByID(id uint) (*Post, error)
	CreatePost(post *Post) (*Post, error)
	UpdatePost(post *Post) (*Post, error)
	DeletePost(id uint) error
	GetPostsByUserID(userID uint) ([]Post, error)
}

type postServiceImpl struct {
	repo PostRepository
}

// CreatePost implements PostService.
func (p *postServiceImpl) CreatePost(post *Post) (*Post, error) {
	post.UserID = CurrentUser.ID
	fmt.Println("CurrentUser.ID", CurrentUser.ID)
	if err := p.repo.CreatePost(post); err != nil {
		return nil, err
	}
	return post, nil
}

// DeletePost implements PostService.
func (p *postServiceImpl) DeletePost(id uint) error {
	panic("unimplemented")
}

// GetAllPosts implements PostService.
func (p *postServiceImpl) GetAllPosts() ([]Post, error) {
	panic("unimplemented")
}

// GetPostByID implements PostService.
func (p *postServiceImpl) GetPostByID(id uint) (*Post, error) {
	post, err := p.repo.GetPostByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// GetPostsByUserID implements PostService.
func (p *postServiceImpl) GetPostsByUserID(userID uint) ([]Post, error) {
	panic("unimplemented")
}

// UpdatePost implements PostService.
func (p *postServiceImpl) UpdatePost(post *Post) (*Post, error) {
	panic("unimplemented")
}

func NewPostService(repo PostRepository) PostService {
	return &postServiceImpl{repo: repo}
}

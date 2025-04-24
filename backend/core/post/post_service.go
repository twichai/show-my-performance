package postCore

import (
	"fmt"
	"mime/multipart"
	s3Adapter "show-my-performance/backend/adapters/s3"
	"show-my-performance/backend/model"
)

type PostService interface {
	GetAllPosts() ([]model.Post, error)
	GetPostByID(id uint) (*model.Post, error)
	CreatePost(post *model.Post, file []*multipart.FileHeader) (*model.Post, error)
	UpdatePost(post *model.Post, file []*multipart.FileHeader) (*model.Post, error)
	DeletePost(id uint) error
	GetPostsByUserID(userID uint) ([]model.Post, error)
}

type postServiceImpl struct {
	repo PostRepository
}

// CreatePost implements PostService.
func (p *postServiceImpl) CreatePost(post *model.Post, file []*multipart.FileHeader) (*model.Post, error) {
	url := ""
	println("File", len(file))
	if len(file) > 0 {
		s3Repository := s3Adapter.NewS3Repository("https://show-my-performance.s3.amazonaws.com/", "us-east-1")
		var err error
		url, err = s3Repository.UploadFile(file, post.Title)
		if err != nil {
			return nil, err
		}
	}
	post.ImageUrl = url
	post.UserID = model.CurrentUser.ID
	fmt.Println("CurrentUser.ID", model.CurrentUser.ID)
	if err := p.repo.CreatePost(post); err != nil {
		return nil, err
	}
	return post, nil
}

// DeletePost implements PostService.
func (p *postServiceImpl) DeletePost(id uint) error {
	if err := p.repo.DeletePost(id, model.CurrentUser.ID); err != nil {
		return err
	}
	return nil

}

// GetAllPosts implements PostService.
func (p *postServiceImpl) GetAllPosts() ([]model.Post, error) {
	post, err := p.repo.GetAllPosts()
	if err != nil {
		return nil, err
	}
	return post, nil
}

// GetPostByID implements PostService.
func (p *postServiceImpl) GetPostByID(id uint) (*model.Post, error) {
	post, err := p.repo.GetPostByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// GetPostsByUserID implements PostService.
func (p *postServiceImpl) GetPostsByUserID(userID uint) ([]model.Post, error) {
	posts, err := p.repo.GetPostsByUserID(userID)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

// UpdatePost implements PostService.
func (p *postServiceImpl) UpdatePost(post *model.Post, file []*multipart.FileHeader) (*model.Post, error) {
	fmt.Println(post.ID)
	oldPost, err := p.repo.GetPostByID(post.ID)
	if err != nil {
		return nil, err
	}
	if oldPost.UserID != model.CurrentUser.ID {
		return nil, fmt.Errorf("you are not authorized to update this post")
	}

	if len(file) > 0 {
		s3Repository := s3Adapter.NewS3Repository("https://show-my-performance.s3.amazonaws.com/", "us-east-1")
		url, err := s3Repository.UploadFile(file, post.Title)
		if err != nil {
			return nil, err
		}
		if url != "" {
			oldPost.ImageUrl = url
		}
	}
	oldPost.Title = post.Title
	oldPost.Content = post.Content
	if err := p.repo.UpdatePost(oldPost); err != nil {
		return nil, err
	}
	return oldPost, nil
}

func NewPostService(repo PostRepository) PostService {
	return &postServiceImpl{repo: repo}
}

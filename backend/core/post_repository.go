package core

type PostRepository interface {
	GetAllPosts() ([]Post, error)
	GetPostByID(id uint) (*Post, error)
	CreatePost(post *Post) error
	UpdatePost(post *Post) error
	DeletePost(id uint, userID uint) error
	GetPostsByUserID(userID uint) ([]Post, error)
}

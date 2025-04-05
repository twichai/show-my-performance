package model

type Reaction struct {
	UserID       uint   `json:"userId"`
	PostID       uint   `json:"postId"`
	ReactionType string `json:"reactionType"` // "like", "dislike"
}

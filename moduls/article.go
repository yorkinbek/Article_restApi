package moduls

import "time"

// Content ...
type Content struct {
	Title string `json:"title" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

// Article ...
type Article struct {
	ID        string     `json:"id"`
	Content              // Promoted fields
	AuthorID  string     `json:"author_id" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

// CreateModelArticle ...
type CreateArticleModel struct {
	Content         // Promoted fields
	AuthorID string `json:"author_id" binding:"required"`
}

// UpdateArticleModel ...
type UpdateArticleModel struct {
	ID      string `json:"id" binding:"required"`
	Content        // Promoted fields
}

// FullArticleModuls ...
type FullArticleModuls struct {
	ID        string     `json:"id"`
	Content              // Promoted fields
	Author    Author     `json:"author"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deletet_at"`
}

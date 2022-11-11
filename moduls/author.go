package moduls

import "time"

// Author ...
type Author struct {
	ID        string     `json:"id"`
	Firstname string     `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"Yorqin"`
	Lastname  string     `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Baqoyev"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

type ArticlesOfAuthor struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"Yorqin"`
	Lastname  string `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Baqoyev"`
	Article   []Article
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}

// CreateAuthorModel ...
type CreateAuthorModel struct {
	Firstname string `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"Yorqin"`
	Lastname  string `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Baqoyev"`
}

// UpdateAuthorModel ...
type UpdateAuthorModel struct {
	ID        string
	Firstname string `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"Yorqin"`
	Lastname  string `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Baqoyev"`
}

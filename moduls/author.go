package moduls

import "time"

// Author ...
type Author struct {
	ID         string     `json:"id"`
	Firstname  string     `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"Yorqin"`
	Lastname   string     `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Baqoyev"`
	Middlename string     `json:"middlename" binding:"required" minLength:"2" maxLength:"50" example:"vich"`
	Fullname   string     `json:"fullname" binding:"required" minLength:"2" maxLength:"70" example:"firstname+lastname+middlename"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"-"`
}

type ArticlesOfAuthor struct {
	ID         string `json:"id"`
	Firstname  string `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"Yorqin"`
	Lastname   string `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Baqoyev"`
	Middlename string `json:"middlename" binding:"required" minLength:"2" maxLength:"50" example:"vich"`
	Fullname   string `json:"fullname" binding:"required" minLength:"2" maxLength:"70" example:"firstname+lastname+middlename"`
	Article    []Article
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"-"`
}

// CreateAuthorModel ...
type CreateAuthorModel struct {
	Firstname  string `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"Yorqin"`
	Lastname   string `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Baqoyev"`
	Middlename string `json:"middlename" binding:"required" minLength:"2" maxLength:"50" example:"vich"`
	Fullname   string `json:"fullname" binding:"required" minLength:"2" maxLength:"70" example:"firstname+lastname+middlename"`
}

// UpdateAuthorModel ...
type UpdateAuthorModel struct {
	ID         string
	Firstname  string `json:"firstname" binding:"required" minLength:"2" maxLength:"50" example:"Yorqin"`
	Lastname   string `json:"lastname" binding:"required" minLength:"2" maxLength:"50" example:"Baqoyev"`
	Middlename string `json:"middlename" binding:"required" minLength:"2" maxLength:"50" example:"vich"`
	Fullname   string `json:"fullname" binding:"required" minLength:"2" maxLength:"70" example:"firstname+lastname+middlename"`
}

package storage

import "github/yorqinbek/CRUD/article/moduls"

// StorageI ...
type StorageI interface {
	AddArticle(id string, entity moduls.CreateArticleModel) error
	GetArticleByID(id string) (moduls.FullArticleModuls, error)
	GetArticleList(offset, limit int, search string) (resp []moduls.Article, err error)
	GetArticlesByAuthorID(id string) ([]moduls.Article, error)
	UpdateArticle(entity moduls.UpdateArticleModel) error
	DeleteArticle(id string) error

	AddAuthor(id string, entity moduls.CreateAuthorModel) error
	GetAuthorByID(id string) (moduls.ArticlesOfAuthor, error)
	GetAuthorList(offset, limit int, search string) (resp []moduls.Author, err error)
	UpdateAuthor(entity moduls.UpdateAuthorModel) error
	DeleteAuthor(id string) error
}

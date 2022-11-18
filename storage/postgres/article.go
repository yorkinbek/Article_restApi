package postgres

import (
	"errors"
	"github/yorqinbek/CRUD/article/moduls"
)

// AddArticle ...
func (stg Postgres) AddArticle(id string, entity moduls.CreateArticleModel) error {
	_, err := stg.GetAuthorByID(entity.AuthorID)
	if err != nil {
		return err
	}

	_, err = stg.db.Exec(`INSERT INTO article 
	(
		id,
		title,
		body,
		author_id
	) VALUES (
		$1,
		$2,
		$3,
		$4
	)`,
		id,
		entity.Title,
		entity.Body,
		entity.AuthorID,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetArticleByID ...
func (stg Postgres) GetArticleByID(id string) (moduls.FullArticleModuls, error) {
	var a moduls.FullArticleModuls

	var tempMiddlename *string
	err := stg.db.QueryRow(`SELECT 
		ar.id,
		ar.title,
		ar.body,
		ar.created_at,
		ar.updated_at,
		ar.deleted_at,
		au.id,
		au.firstname,
		au.lastname,
		au.middlename,
		au.created_at,
		au.updated_at,
		au.deleted_at
    FROM article AS ar JOIN author AS au ON ar.author_id = au.id WHERE ar.id = $1`, id).Scan(
		&a.ID,
		&a.Title,
		&a.Body,
		&a.CreatedAt,
		&a.UpdatedAt,
		&a.DeletedAt,
		&a.Author.ID,
		&a.Author.Firstname,
		&a.Author.Lastname,
		&tempMiddlename,
		&a.Author.CreatedAt,
		&a.Author.UpdatedAt,
		&a.Author.DeletedAt,
	)
	if err != nil {
		return a, err
	}

	if tempMiddlename != nil {
		a.Author.Middlename = *tempMiddlename
	}

	return a, nil
}

// GetArticleList ...
func (stg Postgres) GetArticleList(offset, limit int, search string) (resp []moduls.Article, err error) {
	rows, err := stg.db.Queryx(`SELECT
	id,
	title,
	body,
	author_id,
	created_at,
	updated_at,
	deleted_at 
	FROM article WHERE deleted_at IS NULL AND ((title ILIKE '%' || $1 || '%') OR (body ILIKE '%' || $1 || '%'))
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var a moduls.Article

		err := rows.Scan(
			&a.ID,
			&a.Title,
			&a.Body,
			&a.AuthorID,
			&a.CreatedAt,
			&a.UpdatedAt,
			&a.DeletedAt,
		)
		if err != nil {
			return resp, err
		}
		resp = append(resp, a)
	}

	return resp, err
}

// UpdateArticle ...
func (stg Postgres) UpdateArticle(entity moduls.UpdateArticleModel) error {
	res, err := stg.db.NamedExec("UPDATE article  SET title=:t, body=:b, updated_at=now() WHERE deleted_at IS NULL AND id=:id", map[string]interface{}{
		"id": entity.ID,
		"t":  entity.Title,
		"b":  entity.Body,
	})
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("article not found")
}

// GetArticlesByAuthorID ...
func (stg Postgres) GetArticlesByAuthorID(id string) ([]moduls.Article, error) {
	var resp []moduls.Article
	rows, err := stg.db.Queryx(`SELECT
	id,
	title,
	body,
	author_id,
	created_at,
	updated_at,
	deleted_at 
	FROM article WHERE deleted_at IS NULL AND author_id=$1
	`, id)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		var a moduls.Article

		err := rows.Scan(
			&a.ID,
			&a.Title,
			&a.Body,
			&a.AuthorID,
			&a.CreatedAt,
			&a.UpdatedAt,
			&a.DeletedAt,
		)
		if err != nil {
			return resp, err
		}
		resp = append(resp, a)
	}

	return resp, err
}

// DeleteArticle ...
func (stg Postgres) DeleteArticle(id string) error {
	res, err := stg.db.Exec("UPDATE article  SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL", id)
	if err != nil {
		return err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if n > 0 {
		return nil
	}

	return errors.New("article not found")
}

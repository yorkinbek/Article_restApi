package postgres

import (
	"errors"
	"github/yorqinbek/CRUD/article/moduls"
)

// AddAuthor ...
func (stg Postgres) AddAuthor(id string, entity moduls.CreateAuthorModel) error {
	_, err := stg.db.Exec(`INSERT INTO author 
	(
		id,
		firstname,
		lastname,
		middlename
	) VALUES (
		$1,
		$2,
		$3
	)`,
		id,
		entity.Firstname,
		entity.Lastname,
		entity.Middlename,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetArticleByID ...
func (stg Postgres) GetAuthorByID(id string) (moduls.ArticlesOfAuthor, error) {
	var a moduls.ArticlesOfAuthor
	var tempMiddlename *string
	err := stg.db.QueryRow(`SELECT 
		id,
		firstname,
		lastname,
		middlename,
		created_at,
		updated_at,
		deleted_at
    FROM author WHERE id = $1`, id).Scan(
		&a.ID,
		&a.Firstname,
		&a.Lastname,
		&tempMiddlename,
		&a.CreatedAt,
		&a.UpdatedAt,
		&a.DeletedAt,
	)
	if err != nil {
		return a, err
	}

	if tempMiddlename != nil {
		a.Middlename = *tempMiddlename
	}

	return a, nil
}

// GetAuthorList ...
func (stg Postgres) GetAuthorList(offset, limit int, search string) (resp []moduls.Author, err error) {
	// resp = im.Db.InMemoryAuthorData
	rows, err := stg.db.Queryx(`SELECT
	id,
	firstname,
	lastname,
	middlename,
	created_at,
	updated_at,
	deleted_at
	FROM author WHERE deleted_at IS NULL AND ((firstname ILIKE '%' || $1 || '%') OR (lastname ILIKE '%' || $1 || '%'))
	LIMIT $2
	OFFSET $3
	`, search, limit, offset)
	if err != nil {
		return resp, err
	}
	var tempMiddlename *string
	for rows.Next() {
		var a moduls.Author

		err := rows.Scan(
			&a.ID,
			&a.Firstname,
			&a.Lastname,
			&tempMiddlename,
			&a.CreatedAt,
			&a.UpdatedAt,
			&a.DeletedAt,
		)
		if err != nil {
			return resp, err
		}

		if tempMiddlename != nil {
			a.Middlename = *tempMiddlename
		}
		resp = append(resp, a)
	}
	return resp, err
}

func (stg Postgres) UpdateAuthor(entity moduls.UpdateAuthorModel) error {
	res, err := stg.db.NamedExec("UPDATE author  SET firstname=:f, lastname=:l,middlename=:m, updated_at=now() WHERE deleted_at IS NULL AND id=:id", map[string]interface{}{
		"id": entity.ID,
		"f":  entity.Firstname,
		"l":  entity.Lastname,
		"m":  entity.Middlename,
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

	return errors.New("Authr not found")
}

// DeleteAuthor ...
func (stg Postgres) DeleteAuthor(id string) error {
	res, err := stg.db.Exec("UPDATE author SET deleted_at=now() WHERE id=$1 AND deleted_at IS NULL", id)
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

	return errors.New("Author not found")
}

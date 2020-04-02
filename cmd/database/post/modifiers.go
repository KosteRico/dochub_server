package post

import (
	"checkaem_server/cmd/database"
	obj "checkaem_server/cmd/entities/post"
	"errors"
	"github.com/jackc/pgx"
)

var ErrNotAdmin = errors.New("user doesn't have rights")

func addTags(tx *pgx.Tx, p *obj.Post) error {
	for _, t := range p.TagNames {
		_, err := tx.Exec(addTagQuery, t, p.Id)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	return nil
}

func Insert(p *obj.Post) (*obj.Post, error) {

	tx, err := database.Connection.Begin()

	if err != nil {
		return nil, err
	}

	row := tx.QueryRow(insertQuery, p.Description, p.CreatorUsername, p.Title)

	err = ScanFullPost(row, p)

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	err = addTags(tx, p)

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return p, nil
}

func Delete(id string, username string) (*obj.Post, error) {

	tx, err := database.Connection.Begin()

	if err != nil {
		return nil, err
	}

	p := obj.NewEmpty()

	rows, err := tx.Query(deleteTagQuery, id)

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	var tags []string

	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			_ = tx.Rollback()
			return nil, err
		}
		tags = append(tags, s)
	}

	row := tx.QueryRow(deleteQuery, id)

	err = ScanFullPost(row, p)

	if username != p.CreatorUsername {
		err = tx.Rollback()

		if err != nil {
			return nil, err
		}

		return nil, ErrNotAdmin
	}

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	tx.Commit()

	p.TagNames = tags

	return p, nil
}

func Modify(p *obj.Post) (*obj.Post, error) {

	pBefore, err := Get(p.Id)

	if err != nil {
		return nil, err
	}

	if pBefore.CreatorUsername != p.CreatorUsername {
		return nil, ErrNotAdmin
	}

	if p.Title == "" {
		p.Title = pBefore.Title
	}

	if p.Description == "" {
		p.Description = pBefore.Description
	}

	row := database.Connection.QueryRow(updatePostQuery, p.Description, p.Title, p.Id)

	err = ScanFullPost(row, p)

	if err != nil {
		return nil, err
	}

	return p, nil
}

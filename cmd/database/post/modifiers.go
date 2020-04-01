package post

import (
	"checkaem_server/cmd/database"
	obj "checkaem_server/cmd/entities/post"
	"github.com/jackc/pgx"
)

func addTags(tx *pgx.Tx, tags []string, postId string) error {
	for _, t := range tags {
		_, err := tx.Exec(addTagQuery, t, postId)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	return nil
}

func Insert(description, creatorUsername, title string, tags []string) (*obj.Post, error) {

	tx, err := database.Connection.Begin()

	if err != nil {
		return nil, err
	}

	p := obj.NewEmpty()

	row := tx.QueryRow(insertQuery, description, creatorUsername, title)

	err = ScanFullPost(row, p)

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	err = addTags(tx, tags, p.Id)

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return p, nil
}

func Delete(id string) (*obj.Post, error) {

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

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	tx.Commit()

	p.TagNames = tags

	return p, nil
}

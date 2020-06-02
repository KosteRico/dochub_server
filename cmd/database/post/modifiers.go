package post

import (
	"bytes"
	"checkaem_server/cmd/database"
	obj "checkaem_server/cmd/entities/post"
	"checkaem_server/cmd/searchEngine/indexing"
	"checkaem_server/cmd/tika"
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"io"
	"log"
	"mime/multipart"
	"sync"
)

var ErrNotAdmin = errors.New("user doesn't have rights")

func addTags(tx pgx.Tx, p *obj.Post) error {
	for _, t := range p.TagNames {
		_, err := tx.Exec(context.Background(), addTagQuery, t, p.Id)
		if err != nil {
			_ = tx.Rollback(context.Background())
			return err
		}

		_, err = tx.Exec(context.Background(), updateTagCounter, t)

		if err != nil {
			_ = tx.Rollback(context.Background())
			return err
		}

	}
	return nil
}

func Insert(p *obj.Post) (*obj.Post, error) {

	tx, err := database.Connection.Begin(context.Background())

	if err != nil {
		return nil, err
	}

	row := tx.QueryRow(context.Background(), insertQuery, p.Id, p.Description, p.CreatorUsername, p.Title)

	err = ScanFullPost(p.CreatorUsername, row, p)

	if err != nil {
		_ = tx.Rollback(context.Background())
		return nil, err
	}

	err = addTags(tx, p)

	if err != nil {
		_ = tx.Rollback(context.Background())
		return nil, err
	}

	tx.Commit(context.Background())

	return p, nil
}

func DownloadFile(uuid string) ([]byte, error) {

	var byteArray []byte

	err := database.Connection.QueryRow(context.Background(), selectFileQuery, uuid).Scan(&byteArray)

	if err != nil {
		return nil, err
	}

	return byteArray, nil
}

func UploadFile(file multipart.File, uuid string, size int64) error {
	byteArray := make([]byte, size)

	_, err := file.Read(byteArray)

	if err != nil && err != io.EOF {
		return err
	}

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		r := bytes.NewReader(byteArray)
		text, err := tika.ParseToStr(r)
		if err != nil {
			log.Println(err)
			return
		}

		if tika.IsOCR(text) {
			text, err = tika.ParseToStrOcr(r)
			if err != nil {
				log.Println(err)
				return
			}
		}

		err = indexing.Index(uuid, text)
		if err != nil {
			log.Println(err)
		}
	}(wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		_, err = database.Connection.Exec(context.Background(), insertFileQuery, uuid, byteArray)
		if err != nil {
			log.Println(err)
		}
	}(wg)

	wg.Wait()

	return nil
}

func Delete(id string, username string) (*obj.Post, error) {

	tx, err := database.Connection.Begin(context.Background())

	if err != nil {
		return nil, err
	}

	p := obj.NewEmpty()

	rows, err := tx.Query(context.Background(), deleteTagQuery, id)

	if err != nil {
		_ = tx.Rollback(context.Background())
		return nil, err
	}

	var tags []string

	for rows.Next() {
		var s string
		err = rows.Scan(&s)
		if err != nil {
			_ = tx.Rollback(context.Background())
			return nil, err
		}
		tags = append(tags, s)
	}

	row := tx.QueryRow(context.Background(), deleteQuery, id)

	err = ScanFullPost(p.CreatorUsername, row, p)

	if username != p.CreatorUsername {
		err = tx.Rollback(context.Background())

		if err != nil {
			return nil, err
		}

		return nil, ErrNotAdmin
	}

	if err != nil {
		_ = tx.Rollback(context.Background())
		return nil, err
	}

	tx.Commit(context.Background())

	p.TagNames = tags

	return p, nil
}

func Modify(p *obj.Post) (*obj.Post, error) {

	pBefore, err := Get(p.CreatorUsername, p.Id)

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

	row := database.Connection.QueryRow(context.Background(), updatePostQuery, p.Description, p.Title, p.Id)

	err = ScanFullPost(p.CreatorUsername, row, p)

	if err != nil {
		return nil, err
	}

	return p, nil
}

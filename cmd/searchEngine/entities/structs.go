package entities

import (
	"checkaem_server/cmd/database"
	"context"
	"encoding/json"
)

type DocumentMeta struct {
	Id         string `json:"-"`
	WordsCount int64  `json:"wc"`
}

func NewDocumentMeta(id string, wordsCount int64) *DocumentMeta {
	return &DocumentMeta{
		Id:         id,
		WordsCount: wordsCount,
	}
}

func ExportDocumentMeta(id string) (dm *DocumentMeta, err error) {

	var b []byte

	err = database.Connection.QueryRow(context.Background(),
		selectDocumentMetaQuery, id).Scan(&b)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, dm)

	if err != nil {
		return nil, err
	}

	dm.Id = id

	return
}

func (dm *DocumentMeta) GetInfo(term string) (ii *IndexInfo, err error) {
	var b []byte
	err = database.Connection.QueryRow(context.Background(),
		selectIndexInfoQuery, dm.Id, term).Scan(&b)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, ii)

	if err != nil {
		return nil, err
	}

	return
}

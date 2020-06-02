package entities

import (
	"checkaem_server/cmd/database"
	"context"
	"encoding/json"
)

type IndexInfo struct {
	Id             string `json:"-"`
	PositionsCount int64  `json:"c"`
	Positions      []int  `json:"p"`
}

func NewIndexInfo(id string) *IndexInfo {
	return &IndexInfo{
		Id:             id,
		PositionsCount: 0,
		Positions:      []int{},
	}
}

func (ii *IndexInfo) AppendPosition(val int) {
	ii.Positions = append(ii.Positions, val)
	ii.PositionsCount++
}

func (ii *IndexInfo) UploadToDb(word string) error {
	bytes, err := json.Marshal(ii)

	if err != nil {
		return err
	}

	_, err = database.Connection.Exec(context.Background(), insertIndexInfoQuery, ii.Id, word, bytes)

	return err
}

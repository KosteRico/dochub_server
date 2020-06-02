package entities

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/searchEngine/redisDb"
	"context"
	"math"
)

const (
	K1 = 2.0
	B  = 0.75
)

type Bm25Builder struct {
	DocumentsCounter int64
	DocMeta          *DocumentMeta
}

func NewBm25Builder() *Bm25Builder {
	return &Bm25Builder{
		DocumentsCounter: redisDb.GetDocsCount(),
	}
}

func (bm *Bm25Builder) SetDoc(id string) error {
	err := database.Connection.QueryRow(
		context.Background(),
		selectDocumentMetaQuery, id).Scan(&bm.DocMeta)

	bm.DocMeta.Id = id

	return err
}

func (bm *Bm25Builder) idf(term string) float64 {
	docsWithTerm := redisDb.CountDocuments(term)

	return idfSmooth(bm.DocumentsCounter, docsWithTerm)
}

func (bm *Bm25Builder) tf(term string) (float64, error) {
	var termCount int64

	err := database.Connection.QueryRow(context.Background(),
		selectTermCountQuery,
		bm.DocMeta.Id, term).Scan(&termCount)

	if err != nil {
		return 0, err
	}

	return float64(termCount) / float64(bm.DocMeta.WordsCount), nil
}

func (bm *Bm25Builder) Calc(terms []string) (float64, error) {
	score := 0.0

	avgdl := redisDb.GetAvgDocLength()

	for _, term := range terms {
		tf, err := bm.tf(term)

		if err != nil {
			return 0, err
		}

		idf := bm.idf(term)

		up := idf * (tf * (K1 + 1))

		down := tf + K1*(1-B+B*float64(bm.DocMeta.WordsCount/avgdl))

		score += up / down
	}

	return score, nil
}

func idfSmooth(allDocs, wordDocs int64) float64 {
	down := 1 + wordDocs
	return math.Log(float64(allDocs/down)) + 1
}

func idfCommon(allDocs, wordDocs int64) float64 {
	return math.Log(float64(allDocs / wordDocs))
}

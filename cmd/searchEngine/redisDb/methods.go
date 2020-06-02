package redisDb

import (
	"math/big"
	"strconv"
)

const (
	documentsCounterKey  = "document_counter"
	totalWordsCounterKey = "words_counter"
	averageDocLengthKey  = "avg_doc_length"
)

func GetIntersections(terms []string) []string {
	return termsClient.SInter(terms...).Val()
}

func AddDoc(term string, docId string) bool {
	added := termsClient.SAdd(term, docId)
	return added.Val() == 1
}

func CountDocuments(term string) int64 {
	return termsClient.SCard(term).Val()
}

func IncrementGlobalDocCounter() {
	metaClient.Incr(documentsCounterKey)
}

func SetDocumentMeta(wordsCount int64) {
	IncrementGlobalDocCounter()
	addWordsCount(wordsCount)
	updateAvgDocLength()
}

func addWordsCount(wordsCount int64) {
	metaClient.IncrBy(totalWordsCounterKey, wordsCount)
}

func updateAvgDocLength() {

	twcStr := metaClient.Get(totalWordsCounterKey).Val()
	twcBig, _ := new(big.Int).SetString(twcStr, 10)

	docsCounter := GetDocsCount()

	res := new(big.Int)

	res = res.Div(twcBig, big.NewInt(docsCounter))

	//log.Println(res.Int64())

	metaClient.Set(averageDocLengthKey, res.String(), 0)
}

func GetAvgDocLength() int64 {
	res, _ := metaClient.Get(averageDocLengthKey).Int64()
	return res
}

func GetDocsCount() int64 {
	strVal := metaClient.Get(documentsCounterKey).Val()

	res, _ := strconv.ParseInt(strVal, 10, 64)

	return res
}

package indexing

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/searchEngine/entities"
	"checkaem_server/cmd/searchEngine/redisDb"
	"checkaem_server/cmd/searchEngine/util"
	"context"
	"encoding/json"
	"github.com/bbalet/stopwords"
	"github.com/kljensen/snowball"
	"log"
	"strings"
	"sync"
)

const insertMaxGoroutines int32 = 8000

func handleTerm(wg *sync.WaitGroup, term string, id string) {
	defer wg.Done()
	redisDb.AddDoc(term, id)
}

func isStopword(word string, language string) bool {
	word = stopwords.CleanString(word, language, false)

	return word == "" || word == " "
}

func Index(docId string, text string) error {

	words := util.Field(text)

	l := util.DetectLanguage(text).Iso6391()

	wg := &sync.WaitGroup{}

	indexInfos := make(map[string]*entities.IndexInfo)

	for pos, word := range words {
		word = strings.ToLower(word)
		lang := strings.ToLower(util.DetectLanguage(word).String())

		if isStopword(word, l) {
			continue
		}

		stemmed, err := snowball.Stem(word, lang, false)

		if err != nil {
			return err
		}

		if _, ok := indexInfos[stemmed]; !ok {
			indexInfos[stemmed] = entities.NewIndexInfo(docId)
		}

		indexInfos[stemmed].AppendPosition(pos)

		wg.Add(1)
		go handleTerm(wg, stemmed, docId)
	}

	wg.Wait()

	wg = &sync.WaitGroup{}

	goCounter := util.NewMaxGoroutinesLimiter(insertMaxGoroutines)

	for term, info := range indexInfos {

		goCounter.Wait()

		goCounter.Increment()
		wg.Add(1)
		go func(wg *sync.WaitGroup, term string, info *entities.IndexInfo, counter *util.MaxGoroutinesLimiter) {
			defer wg.Done()

			err := info.UploadToDb(term)

			if err != nil {
				log.Println(err)
			}

			counter.Decrement()
		}(wg, term, info, goCounter)
	}

	wg.Wait()

	log.Println("Final number of goroutines:", goCounter.Counter())

	wordsCount := int64(len(words))

	redisDb.SetDocumentMeta(wordsCount)

	docMeta := entities.NewDocumentMeta(docId, wordsCount)

	b, err := json.Marshal(docMeta)

	if err != nil {
		return err
	}

	_, err = database.Connection.Exec(context.Background(), insertDocMetaQuery, docId, b)

	return err
}

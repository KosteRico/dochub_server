package search

import (
	"checkaem_server/cmd/searchEngine/redisDb"
	"checkaem_server/cmd/searchEngine/util"
	"github.com/kljensen/snowball"
	"strings"
)

func search(query string) (res []string, err error) {
	query = util.CleanStopwords(query)

	words := util.Field(query)

	var terms []string

	for _, word := range words {
		l := strings.ToLower(util.DetectLanguage(word).String())

		stemmed, err := snowball.Stem(word, l, false)

		if err != nil {
			return nil, err
		}

		terms = append(terms, stemmed)
	}

	res = redisDb.GetIntersections(terms)

	return res, nil
}

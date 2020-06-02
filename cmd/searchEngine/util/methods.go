package util

import (
	"github.com/abadojack/whatlanggo"
	"github.com/bbalet/stopwords"
	"github.com/kljensen/snowball"
	"strings"
	"unicode"
)

func DetectLanguage(text string) whatlanggo.Lang {

	options := whatlanggo.Options{
		Whitelist: map[whatlanggo.Lang]bool{
			whatlanggo.Rus: true,
			whatlanggo.Eng: true,
		},
	}

	info := whatlanggo.DetectLangWithOptions(text, options)

	return info
}

func PrepareQuery(query string) ([]string, error) {
	query = CleanStopwords(query)
	words := Field(query)
	return StemAll(words)
}

func CleanStopwords(text string) string {
	l := DetectLanguage(text).Iso6391()

	return stopwords.CleanString(text, l, false)
}

func Field(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
}

func StemAll(words []string) ([]string, error) {
	var res []string
	for _, word := range words {
		l := strings.ToLower(DetectLanguage(word).String())
		word, err := snowball.Stem(word, l, false)
		if err != nil {
			return nil, err
		}
		res = append(res, word)
	}
	return res, nil
}

func GetLang(word string) string {
	c := word[0]
	isLatin := c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
	if isLatin {
		return "english"
	}
	return "russian"
}

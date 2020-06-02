package tika

import (
	"bufio"
	"context"
	"golang.org/x/net/context/ctxhttp"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	ocrLimit = 10.0
)

var ocrLanguages []string

func ParseToStr(file io.ReadSeeker) (string, error) {

	res, err := getClient().Parse(context.Background(), file)

	if err != nil {
		return "", err
	}

	_, err = file.Seek(0, 0)

	return res, err
}

func GetType(file io.ReadSeeker) (string, error) {

	r := bufio.NewReader(file)

	res, err := getClient().Detect(context.Background(), r)

	return res, err
}

func ParseToStrOcr(file io.ReadSeeker) (string, error) {

	client := http.DefaultClient

	request, err := http.NewRequestWithContext(context.Background(), "PUT", server.URL()+"/tika", file)
	if err != nil {
		return "", err
	}

	request.Header["X-Tika-PDFOcrStrategy"] = []string{"ocr_only"}
	request.Header["X-Tika-OCRLanguage"] = []string{strings.Join(ocrLanguages, "+")}
	resp, err := ctxhttp.Do(context.Background(), client, request)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	file.Seek(0, 0)

	body, err := ioutil.ReadAll(resp.Body)

	return string(body), err
}

func IsOCR(text string) bool {
	l := float32(len(text))
	var count float32
	for _, c := range text {
		if c == '\n' {
			count++
			percentage := count / l * 100.0
			if percentage >= ocrLimit {
				return true
			}
		}
	}
	return false
}

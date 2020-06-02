package tika

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"log"
	"testing"
)

func TestParse(t *testing.T) {

	require.Nil(t, Init())

	dir := "./testFiles/"
	byteArray, err := ioutil.ReadFile(dir + "test3.pdf")

	file := bytes.NewReader(byteArray)
	require.Nil(t, err)

	text, err := ParseToStr(file)

	require.Nil(t, err)

	if IsOCR(text) {
		text, err = ParseToStrOcr(file)
		require.Nil(t, err)
	}

	log.Println(text[:300])

}

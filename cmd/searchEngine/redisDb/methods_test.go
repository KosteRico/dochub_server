package redisDb

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestSetDocumentMeta(t *testing.T) {
	godotenv.Load()

	res, err := Init()

	require.Nil(t, err)

	log.Println(res)

	SetDocumentMeta(222)
}

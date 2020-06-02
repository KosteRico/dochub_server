package search

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/searchEngine/redisDb"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
	"time"
)

func TestSearch(t *testing.T) {

	godotenv.Load()

	require.Nil(t, database.Init())

	_, err := redisDb.Init()

	require.Nil(t, err)

	start := time.Now()

	res, err := search("как же он хорош")

	assert.Nil(t, err)

	fmt.Printf("%#v", res)

	fmt.Println("Time left:", time.Since(start))

}

func TestDocumentMeta_GetInfo(t *testing.T) {

	godotenv.Load()

	require.Nil(t, database.Init())

	_, err := redisDb.Init()

	require.Nil(t, err)

	//id := "3a3710cd-f861-43f5-911c-eef78707f444"
	query := "он добьется успеха"

	start := time.Now()

	ids, err := SearchRanking(query)

	assert.Nil(t, err)

	log.Printf("In took %v sec\n", time.Since(start).Seconds())

	for i, id := range ids {
		fmt.Printf("%v. %s\n", i+1, id)
	}

	assert.Nil(t, err)
}

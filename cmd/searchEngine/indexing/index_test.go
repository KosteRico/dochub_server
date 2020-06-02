package indexing

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/searchEngine/redisDb"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
	"time"
)

func TestIndex(t *testing.T) {

	godotenv.Load()

	require.Nil(t, database.Init())

	_, err := redisDb.Init()

	require.Nil(t, err)

	b, err := ioutil.ReadFile("./files/dmrb.txt")

	require.Nil(t, err)

	txt := decodeWin1251(b)

	start := time.Now()

	id, _ := uuid.NewGen().NewV4()

	err = Index(id.String(), txt)

	fmt.Printf("Time: %v\n", time.Since(start))

	assert.Nil(t, err)

	database.Close()
}

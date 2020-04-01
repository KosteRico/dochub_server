package post

import (
	"checkaem_server/cmd/database"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetTags(t *testing.T) {
	require.Nil(t, database.Init())

	tags, err := GetTags("d5892ad2-522c-4855-8ece-2e17bae9f4f3")

	assert.Nil(t, err)

	fmt.Printf("\n%v\n", tags)

	require.Nil(t, database.Close())
}

func TestGetByCreator(t *testing.T) {
	require.Nil(t, database.Init())

	username := "kosterico"

	posts, err := GetByCreator(username)

	assert.Nil(t, err)

	for _, i := range posts {
		assert.Equal(t, username, i.CreatorUsername)
	}

	require.Nil(t, database.Close())
}

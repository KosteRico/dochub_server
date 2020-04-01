package subscription

import (
	"checkaem_server/cmd/database"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var username = "josh365"

func TestGetPosts(t *testing.T) {
	require.Nil(t, database.Init())

	posts, err := GetPosts(username)

	assert.Nil(t, err)

	for _, i := range posts {
		fmt.Printf("\n%v\n", *i)
	}

	require.Nil(t, database.Close())
}

func TestGetTagNames(t *testing.T) {
	require.Nil(t, database.Init())

	tags, err := GetTagNames(username)

	assert.Nil(t, err)

	fmt.Printf("%v\n", tags)

	require.Nil(t, database.Close())
}

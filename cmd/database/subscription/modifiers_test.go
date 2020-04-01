package subscription

import (
	"checkaem_server/cmd/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func checkContains(arr []string, str string) bool {
	for _, t := range arr {
		if t == str {
			return true
		}
	}
	return false
}

func TestInsertDelete(t *testing.T) {
	require.Nil(t, database.Init())

	username := "kosterico"
	tagName := "math"

	err := Insert(username, tagName)
	assert.Nil(t, err)

	tags, err := GetTagNames(username)
	assert.Nil(t, err)

	assert.True(t, checkContains(tags, tagName))

	err = Delete(username, tagName)

	tags, err = GetTagNames(username)
	assert.Nil(t, err)

	assert.False(t, checkContains(tags, tagName))

	require.Nil(t, database.Close())
}

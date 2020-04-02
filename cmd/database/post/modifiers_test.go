package post

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/entities/post"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsertGetDelete(t *testing.T) {
	require.Nil(t, database.Init())

	tags := []string{"math", "podsyp"}

	p := post.New("This text is awesome!!!!", "kosterico", "Some doc", tags)

	pInserted, err := Insert(p)
	assert.Nil(t, err)

	pSelected, err := Get(pInserted.Id)
	assert.Nil(t, err)

	assert.Equal(t, pInserted.Id, pSelected.Id)

	_, err = Delete(pInserted.Id, "kosterico")

	assert.Nil(t, err)

	_, err = Get(pInserted.Id)
	assert.NotNil(t, err)

	assert.Nil(t, database.Close())
}

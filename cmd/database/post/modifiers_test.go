package post

import (
	"checkaem_server/cmd/database"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestInsertGetDelete(t *testing.T) {
	require.Nil(t, database.Init())

	tags := []string{"math", "podsyp"}

	pInserted, err := Insert("This text is awesome!!!!", "kosterico", "Some doc", tags)
	assert.Nil(t, err)

	pSelected, err := Get(pInserted.Id)
	assert.Nil(t, err)

	assert.Equal(t, pInserted.Id, pSelected.Id)

	_, err = Delete(pInserted.Id)

	assert.Nil(t, err)

	_, err = Get(pInserted.Id)
	assert.NotNil(t, err)

	assert.Nil(t, database.Close())
}

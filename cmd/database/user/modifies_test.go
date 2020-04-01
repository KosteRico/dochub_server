package user

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/entities/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertDelete(t *testing.T) {
	assert.Nil(t, database.Init())

	username := "test"

	u, err := user.New(username, "upibaz22")

	assert.Nil(t, err)

	assert.Nil(t, Insert(&u))
	assert.NotNil(t, Insert(&u))

	assert.Nil(t, Delete(username))
	assert.Nil(t, database.Close())
}

package user

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/entities/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {

	assert.NotNil(t, database.InitDB())

	username := "test"

	u, err := user.New(username, "upibaz22")

	assert.NotNil(t, err)

	assert.NotNil(t, Insert(&u))

	assert.NotNil(t, Delete(username))
}

package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitDB(t *testing.T) {

	assert.NotNil(t, InitDB())

	assert.NotNil(t, Conn.Close())
}

package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitDB(t *testing.T) {
	assert.Nil(t, Init())

	Close()
}

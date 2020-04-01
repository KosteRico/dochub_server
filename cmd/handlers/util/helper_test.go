package util

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerateJwtString(t *testing.T) {
	require.Nil(t, godotenv.Load())

	username := "kosterico"

	token, err := GenerateTokenPair(username)
	assert.Nil(t, err)

	fmt.Printf("\nToken: %s\n\n", token)
}

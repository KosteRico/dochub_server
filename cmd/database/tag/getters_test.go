package tag

import (
	"checkaem_server/cmd/database"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGet(t *testing.T) {
	assert.Nil(t, database.Init())

	username := "josh"

	res, err := Get(username, "math")
	assert.Nil(t, err)

	fmt.Printf("%s %d", res.Name, res.Count)
	``
	for _, i := range res.Posts {
		fmt.Printf("\n%v\n", *i)
	}

	database.Close()
}

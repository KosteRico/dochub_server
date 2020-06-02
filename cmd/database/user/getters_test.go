package user

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/entities/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var testUsernames = []string{"johnsmith", "steve", "clark", "albert1999"}
var testPostDescriptions = []string{"Awesome post1", "Awesome post12", "Awesome post123", "Awesome post1234"}
var testPostCreatorUsernames = []string{testUsernames[0], testUsernames[0], testUsernames[1], testUsernames[3]}
var testPostTitle = []string{"title_1", "tt2", "ttt3", "ttul4"}

func createTestUsers() error {
	for _, i := range testUsernames {
		u, err := user.New(i, "1234")

		if err != nil {
			return err
		}

		err = Insert(&u)

		if err != nil {
			return err
		}
	}
	return nil
}

//
//func createTestPosts() {
//
//	for i, _ := range testPostTitle {
//
//	}
//
//}

func deleteTestUsers() error {
	for _, i := range testUsernames {
		err := Delete(i)
		if err != nil {
			return err
		}
	}
	return nil
}

func TestGetAllUsername(t *testing.T) {

	require.Nil(t, database.Init())

	selectedBefore, err := GetAllNames()
	assert.Nil(t, err)

	assert.Nil(t, createTestUsers())

	selectedAfter, err := GetAllNames()

	assert.Nil(t, err)

	assert.Equal(t, len(selectedAfter), len(testUsernames)+len(selectedBefore))

	assert.Nil(t, deleteTestUsers())
}

//
//func TestGetCreatedPosts(t *testing.T) {
//	require.Nil(t, database.Init())
//
//	assert.Nil(t, createTestUsers())
//
//	posts, err := GetCreatedPosts(testUsernames[0])
//
//	assert.Nil(t, err)
//
//	for i, _ := range testPostDescriptions
//}

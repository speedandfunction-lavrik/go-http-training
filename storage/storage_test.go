package storage_test

import (
	"http-go/models"
	"http-go/storage"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("add user", func(t *testing.T) {
		users := new(storage.Users)
		user := models.User{
			UserName: "john",
			Age:      12,
			Email:    "asdasd",
		}

		err := users.Add(&user)
		if err != nil {
			t.Error(err)
		}
	})
}

package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"http-go/models"
	"http-go/storage"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const createUseerURL = "/users/create"

type testUserStorage struct{}

func createPostTestServer(storage storage.Users) http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("/users/create", CreateUser(storage))

	return router
}

func TestCreateUser(t *testing.T) {
	storage := storage.Users{}

	srv := httptest.NewServer(createPostTestServer(storage))
	defer srv.Close()

	user := models.User{
		UserName: "stephan",
		Age:      30,
		Gender:   "male",
	}

	data, err := json.Marshal(user)

	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", srv.URL, createUseerURL), bytes.NewReader(data))

	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println(err)
		t.Error(err)
	}

	fmt.Println(res.StatusCode)
}

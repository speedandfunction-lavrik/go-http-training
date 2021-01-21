package handlers

import (
	"encoding/json"
	"http-go/models"
	"http-go/storage"
	"io/ioutil"
	"net/http"
)

type adder interface {
	Add(user *models.User) error
}

func CreateUser(src storage.Adder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w)

		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		user := models.User{}

		data, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}

		err = json.Unmarshal(data, &user)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}

		err = src.Add(&user)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusCreated)
			w.Write(data)
		}
	}
}

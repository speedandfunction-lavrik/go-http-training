package handlers

import (
	"encoding/json"
	"http-go/models"
	"http-go/storage"
	"io/ioutil"
	"net/http"
)

func UpdateUser(src storage.Updater) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w)

		if r.Method != http.MethodPut {
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

		usr, err := src.Update(&user)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(usr)
		}
	}
}

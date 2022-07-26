package rest

import (
	"encoding/json"
	"net/http"

	"github.com/trite8q1/go-gorm/internal/repository"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	err := repository.AddUser()
	if err != nil {
		panic(err)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	payload, err := repository.GetUser()
	if err != nil {
		panic(err)
	}

	bytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}

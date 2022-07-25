package rest

import (
	"encoding/json"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	payload := map[string]string{
		"status": "ok",
	}
	bytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	w.Write(bytes)
}

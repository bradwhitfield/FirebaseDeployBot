package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type response struct {
	Message string `json:"message"`
}

func F(w http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		log.Panic(err)
		return
	}

	if err := json.NewEncoder(w).Encode(response{Message: string(d)}); err != nil {
		w.WriteHeader(500)
		return
	}
}

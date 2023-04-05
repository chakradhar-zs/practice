package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Value string
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	var result Message
	w.Header().Set("content-type", "application/json")
	if r.Method != http.MethodGet && r.RequestURI == "/ping" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.RequestURI != "/ping" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	result.Value = "pong"
	body, err := json.Marshal(result)
	if err != nil {
		log.Println("")
		return
	}
	w.Write(body)
	return
}

package main

import (
	"encoding/json"
	"net/http"
)

type books struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

var b = []books{
	{Id: 1, Name: "Alice in Wonderland", Author: "Lewis Carrol"},
	{Id: 2, Name: "Pride and Prejudice", Author: "Jane Austen"},
	{Id: 3, Name: "Gulliver’s Travels", Author: "Jonathan Swift"},
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, _ := json.Marshal(b)
	w.Write(body)
}

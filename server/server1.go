package main

import (
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("pong"))
		w.WriteHeader(200)
	} else {
		w.Write([]byte("Expected Get Method"))
		w.WriteHeader(405)
	}

}
func main() {
	http.HandleFunc("/ping", pingHandler)
	http.ListenAndServe(":8080", nil)
}

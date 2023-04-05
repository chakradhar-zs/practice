package main

import "net/http"

func server1(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(200)
		w.Write([]byte("pong"))
	} else {

		w.Write([]byte("Expected Get Method"))
		w.WriteHeader(405)
	}
}

func main() {
	http.HandleFunc("/ping", server1)
	http.ListenAndServe(":8080", nil)
}

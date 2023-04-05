package main

import (
	"encoding/json"
	"net/http"
)

type Product struct {
	Id          int `json:"id"`
	Name        string
	Description string
	Price       float64
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{1, "apple", "a fresh fruit", 20},
		{2, "orange", "a citrus fruit", 30},
	}
	pro, _ := json.Marshal(products)

	if r.Method == "GET" {

		w.WriteHeader(200)
		w.Write(pro)

	} else {
		w.WriteHeader(405)
		return
	}

}
func main() {
	http.HandleFunc("/product", ProductHandler)
	http.ListenAndServe(":8080", nil)
}

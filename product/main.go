package main

import (
	"encoding/json"
	"io"
	"net/http"
)

//type Product struct {
//	Id          int ``
//	Name        string
//	Description string
//	Price       float64
//}
//
//var products = []Product{
//	{Id: 1, Name: "Product 1", Description: "Description 1", Price: 1.99},
//	{Id: 2, Name: "Product 2", Description: "Description 2", Price: 2.99},
//	{Id: 3, Name: "Product 3", Description: "Description 3", Price: 3.99},
//}

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid Body"))
			return
		}

		//e := json.Unmarshal(body, &prod)
		//if e != nil {
		//	w.WriteHeader(http.StatusBadRequest)
		//	w.Write([]byte("Invalid Body"))
		//	return
		//}
		//for p := range products {
		//	if products[p].Id == prod.Id {
		//		w.WriteHeader(http.StatusBadRequest)
		//		return
		//
		//	}
		//
		//}
		//w.WriteHeader(http.StatusCreated)

		w.Write(body)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func GetProducts(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		prod, err := json.Marshal(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(""))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(prod)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(""))
		return
	}
}

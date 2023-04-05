package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"path"
	"productcreateandretrieve/db"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) int64 {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return 0
	}

	var data db.Data

	p, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(p, &data)
	if err != nil {
		log.Println(err)
	}
	return db.CreateProduct(data, &data.Pro)

}

func getByID(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(""))
	}

	str := path.Base(r.URL.String())
	var p1 db.Product
	pid, err := strconv.Atoi(str)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(""))
	}

	prod, err := db.RetieveProduct(pid, p1)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(""))
	}

	body, err := json.Marshal(prod)
	if err != nil {
		log.Println("Error while marshalling ", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

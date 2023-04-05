package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"path"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Price       int    `json:"price"`
	Quantity    int    `json:"q"`
	Category    string `json:"category"`
	Brand_id    int    `json:"brandId"`
	Status      string `json:"status"`
}

func connectDatabase() *sql.DB {
	conn := "root:password@tcp(localhost:3306)/zopstore"

	db, err := sql.Open("mysql", conn)

	if err != nil {
		log.Println("Failed to connect to database", err)

	}

	if err := db.Ping(); err != nil {
		log.Println("Failed to connect to database", err)

	}

	return db
}

func getByID(w http.ResponseWriter, r *http.Request) {
	db := connectDatabase()

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(""))
	}

	str := path.Base(r.URL.String())
	var p1 Product
	pid, err := strconv.Atoi(str)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(""))
	}

	row := db.QueryRow("select * from products where id = ?", pid)

	err = row.Scan(&p1.Id, &p1.Name, &p1.Description, &p1.Price, &p1.Quantity, &p1.Category, &p1.Brand_id, &p1.Status)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(""))
	}

	body, err := json.Marshal(p1)
	if err != nil {
		log.Println("Error while marshalling ", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

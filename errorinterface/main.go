package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type DBError struct {
	Err error `json:"err"`
}

type Result struct {
	val1 Product
	val2 error
}
type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Price       int    `json:"price"`
	Quantity    int    `json:"q"`
	Category    string `json:"category"`
	Brand_id    int    `json:"brandid"`
	Status      string `json:"status"`
}

func (e DBError) Error() string {
	return fmt.Sprintf("Error: %v", e.Err)
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

func retrieveById(id int, db *sql.DB) (Product, error) {
	db = connectDatabase()

	var p1 Product

	row := db.QueryRow("select * from products where p_id=?", id)

	if row.Err() != nil {
		return Product{}, DBError{Err: row.Err()}
	}
	err := row.Scan(&p1.Id, &p1.Name, &p1.Description, &p1.Price, &p1.Quantity, &p1.Category, &p1.Brand_id, &p1.Status)

	if err != nil {
		return Product{}, DBError{Err: err}
	}

	return p1, DBError{}
}

func RetrieveProduct(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(""))
	}

	var db *sql.DB
	var res Result
	str := path.Base(r.URL.String())
	pid, err := strconv.Atoi(str)

	prod, dberr := retrieveById(pid, db)
	if dberr != nil {
		res.val2 = dberr
	}

	res.val1 = prod
	body, err := json.Marshal(res)
	if err != nil {
		log.Println("Error while marshalling ", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)

}

func main() {
	http.HandleFunc("/product/1", RetrieveProduct)

	http.ListenAndServe(":3000", nil)
}

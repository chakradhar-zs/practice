package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

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

var db *sql.DB

//CreateProduct takes an product details and brand name as input and returns no of rows inserted
func CreateProduct(p Product, brand string) int64 {

	conn := "root:password@tcp(localhost:3306)/zopstore"

	db, err := sql.Open("mysql", conn)

	if err != nil {
		log.Println("Failed to connect to database", err)

	}

	if err := db.Ping(); err != nil {
		log.Println("Failed to connect to database", err)

	}

	res, err := db.Exec("select exists(select * from brands where name=?)", brand)

	isExist, _ := res.RowsAffected()
	if isExist == 0 {
		_, err := db.Exec("insert into brands(id,name) values (?,?)", p.Id, brand)
		if err != nil {
			log.Println(err)
		}
	}

	res, err = db.Exec("insert into products(p_id,name,description,price,quantity,category,Brand_id,status)values(?,?,?,?,?,?,?,?)", p.Id, p.Name, p.Description, p.Price, p.Quantity, p.Category, p.Brand_id, p.Status)

	if err != nil {
		log.Println("Error while inserting, err :", err)
	}
	val, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	return val
}

//RetrieveProducts  takes id as input  returns the product with id that matches the given id and brand name
func RetrieveProducts(val int) (Product, string) {
	conn := "root:password@tcp(localhost:3306)/ZopStore"
	db, err := sql.Open("mysql", conn)

	if err != nil {
		log.Println("Failed to connect to database", err)

	}

	if err := db.Ping(); err != nil {
		log.Println("Failed to connect to database", err)

	}
	var p1 Product
	var bname string
	res := db.QueryRow("select * from products where p_id=? ", val)

	err = res.Scan(&p1.Id, &p1.Name, &p1.Description, &p1.Price, &p1.Quantity, &p1.Category, &p1.Brand_id, &p1.Status)

	res = db.QueryRow("select name from brands where id=?", p1.Brand_id)

	err = res.Scan(&bname)

	if err != nil {
		log.Println(err)
	}

	return p1, bname

}

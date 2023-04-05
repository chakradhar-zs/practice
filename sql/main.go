package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	conn := "root:password@tcp(localhost:3306)/test"

	db, err := sql.Open("mysql", conn)

	if err != nil {
		log.Println("Failed to connect to database", err)
		return
	}

	if err := db.Ping(); err != nil {
		log.Println("Failed to connect to database", err)
		return
	}
	fmt.Println(db.Stats())

	const DropProductsTable = "drop table products"

	const CreateProductsTable = `create table if not exists products(
    id int primary key AUTO_INCREMENT ,
    pname varchar(50) not null ,
    brand varchar(100) null
    )`

	//if _, err = db.Exec(DropProductsTable); err != nil {
	//	log.Println("Failed to delete the table ,err:", err)
	//}
	//
	//if _, err = db.Exec(CreateProductsTable); err != nil {
	//	log.Println("Failed to create the table ,err:", err)
	//}

	resp, error := db.Exec("insert into products(pname,brand) values ('ghee','Amul')")

	if error != nil {
		log.Println("Failed to insert into the table ,err:", error)
		return
	} else {
		id, _ := resp.LastInsertId()
		log.Println("Product Amul Ghee has been created, id of the product :", id)
	}

	row := db.QueryRow("select * from products")
	var id int
	var name string
	var brand string

	if err := row.Scan(&id, &name, &brand); err != nil {
		if err == sql.ErrNoRows {
			log.Println("No products exist")
		}
		log.Println(err)
	}

	log.Println("Product one is :", id, name, brand)
}

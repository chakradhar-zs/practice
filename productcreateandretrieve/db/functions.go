package db

import (
	"database/sql"
	"log"
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

type Data struct {
	Pro   Product
	Brand string
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
func CreateProduct(d1 Data, pro *Product) int64 {
	db := connectDatabase()

	res, err := db.Exec("select id from brands where name =?", d1.Brand)

	isExist, err := res.RowsAffected()
	if isExist == 0 {
		_, err = db.Exec("insert into brands values (?,?)", pro.Brand_id, d1.Brand)
		if err != nil {
			log.Println(err)
		}
	}

	res, err = db.Exec("insert into products values(?,?,?,?,?,?,?,?)", pro.Id, pro.Name, pro.Description, pro.Price, pro.Quantity, pro.Category, pro.Brand_id, pro.Status)
	if err != nil {
		log.Println(err)
	}
	row, _ := res.RowsAffected()
	return row
}

func RetieveProduct(pid int, p1 Product) (Product, error) {
	db := connectDatabase()

	row := db.QueryRow("select * from products where id = ?", pid)
	err := row.Scan(&p1.Id, &p1.Name, &p1.Description, &p1.Price, &p1.Quantity, &p1.Category, &p1.Brand_id, &p1.Status)

	return p1, err
}

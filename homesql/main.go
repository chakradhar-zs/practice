package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	conn := "root:password@tcp(localhost:3306)/ZopStore"

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

	_, err = db.Exec("create table Brands(BrandId int,bname varchar(50) unique ,primary key (BrandId))")

	_, err = db.Exec("create table Products(Id int,pname varchar(50),description varchar(100),price decimal,quantity int,category varchar(50),BrandId int,primary key (Id),foreign key (BrandId)references Brands(BrandId))")

}

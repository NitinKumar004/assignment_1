package main

import (
	"database/sql"
	"fmt"
	"log"
)

func Calldatabase() *sql.DB {

	driver_link := "root:root123@tcp(127.0.0.1:3306)/test_db"
	db, err := sql.Open("mysql", driver_link)
	if err != nil {
		log.Fatal("error to connect the database", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("database is unreachable", err)

	}
	fmt.Println("successfully connected to the database")
	return db
}

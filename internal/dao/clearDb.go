package dao

import (
	"database/sql"
	"log"
	"naval/resources"
)

func ClearDb() {
	db, err := sql.Open("postgres", resources.ConnectDb)
	if err != nil {
		log.Println("Error connecting to dao")
		return
	}
	defer db.Close()
	_, err = db.Exec("drop table info_city")
	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec("DELETE FROM orders")
	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec("DELETE FROM cities")
	if err != nil {
		log.Println(err)
	}
	_, err = db.Exec("DELETE FROM items")
	if err != nil {
		log.Println(err)
	}
}

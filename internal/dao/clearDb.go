package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func ClearDb() {
	//db, err := sql.Open("postgres", resources.ConnectDb)
	db, err := sql.Open("sqlite3", "file:identifier.sqlite")
	if err != nil {
		log.Println("Error connecting to dao")
		return
	}
	defer db.Close()

	version := ""
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)

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

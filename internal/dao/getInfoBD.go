package dao

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func GerInfoDB(request string) *sql.Rows {
	db, err := sql.Open("sqlite3", "file:locked.sqlite")
	if err != nil {
		log.Println("Error connecting to dao")
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM info_city where item_name = $1 ", request)
	if err != nil {
		log.Println(err)
	}
	return rows
}

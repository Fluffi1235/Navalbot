package dao

import (
	"database/sql"
	"log"
	"naval/resources"
)

func GerInfoDB(request string) *sql.Rows {
	db, err := sql.Open("postgres", resources.ConnectDb)
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

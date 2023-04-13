package dao

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func CreatTableInfo_city() {
	db, err := sql.Open("sqlite3", "file:locked.sqlite")
	if err != nil {
		log.Println("Error connecting to dao")
		return
	}
	defer db.Close()

	query := `create table info_city as select c.name as city_name,
	 i.name as item_name, o.quantity, o.price from orders o left join cities c on o.city_id = c.id_city left join items i on o.item_id = i.id_item`
	_, err = db.Exec(query)
	if err != nil {
		log.Println(err)
		return
	}
}

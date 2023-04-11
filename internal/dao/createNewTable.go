package dao

import (
	"database/sql"
	"log"
	"naval/resources"
)

func CreatTableInfo_city() {
	db, err := sql.Open("postgres", resources.ConnectDb)
	if err != nil {
		log.Println("Error connecting to dao")
		return
	}
	defer db.Close()
	_, err = db.Exec("Create table info_city as select c.name as city_name, i.name as item_name, o.quantity, o.price from orders o left join cities c on o.city_id = c.id_city left join items i on o.item_id = i.id_item")
	if err != nil {
		log.Println(err)
		return
	}
}

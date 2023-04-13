package dao

import (
	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"naval/internal/model"
	"strings"
)

func SavePb(town model.Pb) {
	db, err := sql.Open("sqlite3", "file:locked.sqlite")
	if err != nil {
		log.Println("Error connecting to dao")
		return
	}
	defer db.Close()

}

func SaveItems(items model.Items) {
	db, err := sql.Open("sqlite3", "file:locked.sqlite")
	if err != nil {
		log.Println("Error connecting to dao")
		return
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO items(id_item, name) VALUES($1, $2)", items.Id, strings.ToLower(items.NameItem))
	if err != nil {
		log.Println("Error inserting into dao")
		return
	}
}

func SaveCitiInfo(port model.Port) {
	db, err := sql.Open("sqlite3", "file:locked.sqlite")
	if err != nil {
		log.Println("Error connecting to dao")
		return
	}
	defer db.Close()
	for _, value := range port.Invt {
		_, err = db.Exec("INSERT INTO orders(city_id, item_id, quantity, price) VALUES($1, $2, $3, $4)",
			port.Id, value.Id, value.BuyQuantityItem, value.BuyPriceItem)
		if err != nil {
			log.Println("Error inserting into dao")
			return
		}
	}
}

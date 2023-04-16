package repository

import (
	"database/sql"
	"log"
	"naval/internal/model"
	"strings"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) NavalRepo {
	return &Repository{
		db: db,
	}
}

func (r Repository) CreatTableInfo_city() error {
	_, err := r.db.Exec("Create table info_city as select c.name as city_name, i.name as item_name, o.quantity, o.price from orders o left join cities c on o.city_id = c.id_city left join items i on o.item_id = i.id_item")
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r Repository) ClearDB() error {
	_, err := r.db.Exec("drop table info_city")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = r.db.Exec("DELETE FROM orders")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = r.db.Exec("DELETE FROM cities")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = r.db.Exec("DELETE FROM items")
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r Repository) SavePb(town model.Pb) error {
	_, err := r.db.Exec("INSERT INTO cities(id_city, name) VALUES($1, $2)", town.Id, town.NameTown)
	if err != nil {
		log.Println("Error inserting into dao")
		return err
	}
	return nil
}

func (r Repository) SaveItems(items model.Items) error {
	_, err := r.db.Exec("INSERT INTO items(id_item, name) VALUES($1, $2)", items.Id, strings.ToLower(items.NameItem))
	if err != nil {
		log.Println("Error inserting into dao")
		return err
	}
	return nil
}

func (r Repository) SaveCitiInfo(port model.Port) error {
	for _, value := range port.Invt {
		_, err := r.db.Exec("INSERT INTO orders(city_id, item_id, quantity, price) VALUES($1, $2, $3, $4)",
			port.Id, value.Id, value.BuyQuantityItem, value.BuyPriceItem)
		if err != nil {
			log.Println("Error inserting into dao")
			return err
		}
	}
	return nil
}

func (r Repository) GetInfoDB(request string) (*sql.Rows, error) {
	rows, err := r.db.Query("SELECT * FROM info_city where item_name = $1", request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return rows, nil
}

package repository

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"naval/internal/model"
	"strings"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) NavalRepo {
	return &repository{
		db: db,
	}
}

func (r repository) CreatTableInfoCity() error {
	query := `create table IF NOT EXISTS cities (
     id_city   int primary key,
     name   varchar (50)
);

create table IF NOT EXISTS items (
     id_item    int primary key ,
     name  varchar(50)
);

create table IF NOT EXISTS orders (
    id integer primary key AUTOINCREMENT ,
     city_id   int,
     item_id   int,
     quantity int,
     price    int
)`
	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) ClearDB() error {
	_, err := r.db.Exec("DELETE FROM orders")
	if err != nil {
		return err
	}
	_, err = r.db.Exec("DELETE FROM cities")
	if err != nil {
		return err
	}

	_, err = r.db.Exec("DELETE FROM items")
	if err != nil {
		return err
	}

	return nil
}

func (r repository) SavePb(town model.Pb) error {
	_, err := r.db.Exec("INSERT INTO cities(id_city, name) VALUES($1, $2)", town.Id, town.NameTown)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) SaveItems(items model.Items) error {
	_, err := r.db.Exec("INSERT INTO items(id_item, name) VALUES($1, $2)", items.Id, strings.ToLower(items.NameItem))
	if err != nil {
		return err
	}

	return nil
}

func (r repository) SaveCitiInfo(port model.Port) error {
	for _, value := range port.Invt {
		_, err := r.db.Exec("INSERT INTO orders(city_id, item_id, quantity, price) VALUES($1, $2, $3, $4)",
			port.Id, value.Id, value.BuyQuantityItem, value.BuyPriceItem)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r repository) GerInfoDB(request string) (*sql.Rows, error) {
	rows, err := r.db.Query("SELECT * FROM info_city where item_name = $1 ", request)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

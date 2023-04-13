package repository

import (
	"database/sql"
	"naval/internal/model"
)

type NavalRepo interface {
	CreatTableInfoCity() error
	ClearDB() error
	SavePb(town model.Pb) error
	SaveItems(items model.Items) error
	SaveCitiInfo(port model.Port) error

	GerInfoDB(request string) (*sql.Rows, error)
}

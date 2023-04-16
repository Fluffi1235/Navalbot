package repository

import (
	"database/sql"
	"naval/internal/model"
)

type NavalRepo interface {
	CreatTableInfo_city() error
	ClearDB() error
	SavePb(town model.Pb) error
	SaveItems(items model.Items) error
	SaveCitiInfo(port model.Port) error
	GetInfoDB(request string) (*sql.Rows, error)
}

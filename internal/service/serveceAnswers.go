package service

import (
	"log"
	"naval/internal/dao"
	"naval/internal/model"
)

func GerInfoDB(request string) []string {
	answer := make([]string, 0, 0)
	infocity := model.Answer{}
	rows := dao.GerInfoDB(request)
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&infocity.City, &infocity.Item, &infocity.Quantity, &infocity.Price); err != nil {
			log.Println(err)
		}
		answer = append(answer, "Город: "+infocity.City+"; Количество: "+infocity.Quantity+"; Цена: "+infocity.Price+"\n")
	}
	return answer
}

func ClearDb() {
	dao.ClearDb()
}

func SavePb(town model.Pb) {
	dao.SavePb(town)
}

func SaveItems(items model.Items) {
	dao.SaveItems(items)
}

func SaveCitiInfo(port model.Port) {
	dao.SaveCitiInfo(port)
}

func CreatNewTable() {
	dao.CreatTableInfo_city()
}

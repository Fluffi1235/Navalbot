package service

import (
	"log"
	"naval/internal/model"
	"naval/internal/repository"
)

func GerInfoDB(request string, repo repository.NavalRepo) ([]string, error) {
	answer := make([]string, 0, 0)
	infocity := model.Answer{}
	rows, err := repo.GerInfoDB(request)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&infocity.City, &infocity.Item, &infocity.Quantity, &infocity.Price); err != nil {
			log.Println(err)
		}
		answer = append(answer, "Город: "+infocity.City+"; Количество: "+infocity.Quantity+"; Цена: "+infocity.Price+"\n")
	}

	return answer, nil
}

package service

import (
	"log"
	"naval/internal/model"
	"naval/internal/repository"
)

func GerInfoDB(request string, repo repository.NavalRepo) []string {
	answermas := make([]string, 0, 0)
	var counter int
	var answer string
	infocity := model.Answer{}
	rows, _ := repo.GetInfoDB(request)
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&infocity.City, &infocity.Item, &infocity.Quantity, &infocity.Price); err != nil {
			log.Println(err)
		}
		if counter == 25 {
			answermas = append(answermas, answer)
			answer = ""
			counter = 0
		}
		answer = answer + "Город: " + infocity.City + "; Количество: " + infocity.Quantity + "; Цена: " + infocity.Price + "\n"
		counter++
	}
	if len(answermas) == 0 {
		answermas = append(answermas, answer)
	}
	return answermas
}

package parse

import (
	"encoding/json"
	"log"
	"naval/internal/model"
	"naval/internal/repository"
	"net/http"
)

type Parser struct {
	repo repository.NavalRepo
}

func New(repo repository.NavalRepo) *Parser {
	return &Parser{
		repo: repo,
	}
}

func (p *Parser) ParceCity() {
	resp, err := http.Get(model.Pblink)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	pb := make([]model.Pb, 0, 0)

	err = json.NewDecoder(resp.Body).Decode(&pb)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range pb {
		err = p.repo.SavePb(v)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (p *Parser) Port() {
	resp, err := http.Get(model.Portslink)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	port := make([]model.Port, 0, 0)
	err = json.NewDecoder(resp.Body).Decode(&port)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range port {
		err = p.repo.SaveCitiInfo(v)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func (p *Parser) Items() {
	resp, err := http.Get(model.Itemslink)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	items := make([]model.Items, 0, 0)
	err = json.NewDecoder(resp.Body).Decode(&items)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range items {
		err = p.repo.SaveItems(v)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

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
	decoder := json.NewDecoder(resp.Body)
	pb := make([]model.Pb, 0, 0)
	err = decoder.Decode(&pb)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range pb {
		p.repo.SavePb(v)
	}
}

func (p *Parser) Port() {
	resp, err := http.Get(model.Portslink)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	port := make([]model.Port, 0, 0)
	err = decoder.Decode(&port)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range port {
		p.repo.SaveCitiInfo(v)
	}
}

func (p *Parser) Items() {
	resp, err := http.Get(model.Itemslink)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	items := make([]model.Items, 0, 0)
	err = decoder.Decode(&items)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range items {
		p.repo.SaveItems(v)
	}
}

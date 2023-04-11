package parse

import (
	"encoding/json"
	"log"
	"naval/internal/model"
	"naval/internal/service"
	"net/http"
)

func ParceCity() {
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
		service.SavePb(v)
	}
}

func Port() {
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
		service.SaveCitiInfo(v)
	}
	service.CreatNewTable()
}

func Items() {
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
		service.SaveItems(v)
	}
}

package model

const (
	Itemslink string = "https://na-map.netlify.app/data/eu1-items.json"
	Portslink string = "https://na-map.netlify.app/data/eu1-ports.json"
	Pblink    string = "https://na-map.netlify.app/data/eu1-pb.json"
)

type Items struct {
	Id       int    `json:"id"`
	NameItem string `json:"name"`
}

type Port struct {
	Id   int         `json:"id"`
	Invt []Inventory `json:"inventory"`
}

type Inventory struct {
	Id              int `json:"id"`
	BuyQuantityItem int `json:"buyQuantity"`
	BuyPriceItem    int `json:"buyPrice"`
}

type Pb struct {
	Id       int    `json:"id"`
	NameTown string `json:"name"`
}

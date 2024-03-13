package models

import "github.com/google/uuid"

type Card struct {
	ID              uuid.UUID  `json:"id"`
	Name            string     `json:"name"`
	FlavorName      string     `json:"flavor_name"`
	CardDescription string     `json:"oracle_text"`
	ManaCost        string     `json:"mana_cost"`
	Legalities      Legalities `json:"legalities"`
	ImageUris       ImageUris  `json:"image_uris"`
}

type Legalities struct {
	Standard        string `json:"standard"`
	Future          string `json:"future"`
	Historic        string `json:"historic"`
	Timeless        string `json:"timeless"`
	Gladiator       string `json:"gladiator"`
	Pioneer         string `json:"pioneer"`
	Explorer        string `json:"explorer"`
	Modern          string `json:"modern"`
	Legacy          string `json:"legacy"`
	Pauper          string `json:"pauper"`
	Vintage         string `json:"vintage"`
	Penny           string `json:"penny"`
	Commander       string `json:"commander"`
	Oathbreaker     string `json:"oathbreaker"`
	Standardbrawl   string `json:"standardbrawl"`
	Brawl           string `json:"brawl"`
	Alchemy         string `json:"alchemy"`
	Paupercommander string `json:"paupercommander"`
	Duel            string `json:"duel"`
	Oldschool       string `json:"oldschool"`
	Premodern       string `json:"premodern"`
	Predh           string `json:"predh"`
}

type ImageUris struct {
	ArtCrop    string `json:"art_crop`
	BorderCrop string `json:"border_crop`
	Large      string `json:"large`
	Normal     string `json:"normal`
	Png        string `json:"png`
	Small      string `json:"small`
}

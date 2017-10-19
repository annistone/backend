package models

type Link struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
	Link string `json:"link"`
}

type Location struct {
	Id  int     `json:"id"`
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

type MapItem struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Location    Location `json:"location"`
	Links       []Link   `json:"links"`
}

type Achievement struct {
	Id          int    `json:"id"`
	Text        string `json:"text"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type User struct {
	Id           int           `json:"id"`
	Name         string        `json:"name"`
	Last_name    string        `json:"last_name"`
	Rating       int           `json:"rating"`
	Image        string        `json:"image"`
	Achievements []Achievement `json:"achievements"`
}

type Bonus struct {
	Price       int         `json:"price"`
	Achievement Achievement `json:"achievement"`
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

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

type Sight struct {
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

func sendMapItems(w http.ResponseWriter, r *http.Request) {
	palace := []Sight{
		{
			Id:          1,
			Name:        "Palace",
			Image:       "https://www.riu.com/es/binaris/piscina-riu-palace-las-americas-2016_tcm49-157129.jpg",
			Description: "olololo",
			Location: Location{
				Id:  1,
				Lat: 59.956377,
				Lng: 30.309408,
			},
			Links: []Link{
				{
					Id:   1,
					Name: "BigPalace",
					Type: 4,
					Link: "palace.org",
				},
				{
					Id:   1,
					Name: "palace",
					Type: 4,
					Link: "palace.org",
				},
			},
		},
		{
			Id:          1,
			Name:        "Palace2",
			Description: "Palace was builded in 1332.",
			Image:       "https://taj.tajhotels.com/content/dam/luxury/hotels/Umaid_Bhavan/images/16x7/33397215.jpg",
			Location: Location{
				Id:  1,
				Lat: 59.956333,
				Lng: 30.309487,
			},
			Links: []Link{
				{
					Id:   1,
					Name: "OldPalace",
					Type: 4,
					Link: "palace.org",
				},
				{
					Id:   1,
					Name: "palace",
					Type: 4,
					Link: "palace.org",
				},
			},
		},
	}

	b, err := json.Marshal(palace)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Fprintf(w, string(b))
}

func sendUsrInfo(w http.ResponseWriter, r *http.Request) {
	user := []User{
		{
			Id:        1,
			Name:      "Tanya",
			Last_name: "Tanya",
			Rating:    45,
			Image:     "http://i.imgur.com/Rl8Upz0.jpg",
			Achievements: []Achievement{
				{
					Id:          0,
					Text:        "Вы дошли!",
					Description: "Вы дошли до этого места. Возьмите пряник!",
					Image:       "http://host/image.jpg",
				},
				{
					Id:          0,
					Text:        "Вы дошли!",
					Description: "Вы дошли до этого места. Возьмите пряник!",
					Image:       "http://host/image.jpg",
				},
				{
					Id:          0,
					Text:        "Вы дошли!",
					Description: "Вы дошли до этого места. Возьмите пряник!",
					Image:       "http://host/image.jpg",
				},
			},
		},
	}

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Fprintf(w, string(b))
}

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/api/getMapItems", sendMapItems)
	http.HandleFunc("/api/getUsrInfo", sendUsrInfo)

	log.Fatal(http.ListenAndServe(":"+port, nil))

}

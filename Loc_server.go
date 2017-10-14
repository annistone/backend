package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Link struct {
	Id   int
	Name string
	Type int
	Link string
}

type Location struct {
	Id  int
	Lat float32
	Lng float32
}

type Sight struct {
	Id          int
	Name        string
	Description string
	Location    Location
	Links       []Link
}

type Achievement struct {
	Id          int
	Text        string
	Description string
	Image       string
}

type User struct {
	Id           int
	Name         string
	Last_name    string
	Rating       int
	Image        string
	Achievements []Achievement
}

func sendMapItems(w http.ResponseWriter, r *http.Request) {
	palace := []Sight{
		{
			Id:          1,
			Name:        "Palace",
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

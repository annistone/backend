package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	b, err := ioutil.ReadFile("base.json") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Fprintf(w, str)
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

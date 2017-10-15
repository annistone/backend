package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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

type Prize struct {
	Price       int         `json:"price"`
	Achievement Achievement `json:"achievement"`
}

var user User

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

func sendUserInfo(w http.ResponseWriter, r *http.Request) {

	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Fprintf(w, string(b))
}

func sendPrize(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	achieveIdStr := r.FormValue("id")
	achieveId, err := strconv.Atoi(achieveIdStr)
	if err != nil {
		fmt.Println("error:", err)
	}

	price := 100

	for _, userAchieveId := range user.Achievements {
		if userAchieveId.Id == achieveId {
			price = 0
		}
	}

	newAchievement := Achievement{
		Id:          achieveId,
		Text:        "Вы дошли!",
		Description: "Вы дошли до этого места. Возьмите пряник!",
		Image:       "https://i.pinimg.com/236x/a5/76/99/a57699849fb0d8f69c8e4016457b5c66--job-well-done-quotes-congratulations-quotes.jpg",
	}

	if price != 0 {
		user.Achievements = append(user.Achievements, newAchievement)
	}

	user.Rating += price

	prize := Prize{
		Price:       price,
		Achievement: newAchievement,
	}

	b, err := json.Marshal(prize)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Fprintf(w, string(b))
}

func main() {

	user = User{
		Id:           1,
		Name:         "Tanya",
		Last_name:    "Tanya",
		Rating:       45,
		Image:        "http://i.imgur.com/Rl8Upz0.jpg",
		Achievements: []Achievement{},
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/api/getMapItems", sendMapItems)
	http.HandleFunc("/api/getUserInfo", sendUserInfo)
	http.HandleFunc("/api/sendPosition", sendPrize)

	log.Fatal(http.ListenAndServe(":"+port, nil))

}

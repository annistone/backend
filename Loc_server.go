package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	Id       int
	Name     string
	Location Location
	Links    []Link
}

func sendData(w http.ResponseWriter, r *http.Request) {
	palace := Sight{
		Id:   1,
		Name: "Palace",
		Location: Location{
			Id:  1,
			Lat: 0.23,
			Lng: 9.3,
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
	}

	b, err := json.Marshal(palace)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Fprintf(w, string(b))
}

func main() {
	http.HandleFunc("/getData", sendData)

	log.Fatal(http.ListenAndServe(":8081", nil))

}

package main

import (
	"Loc_server/models"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

// MkRegisterHandler function
func MkRegisterHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendMapItems(w, r, db)
	}
}

func sendMapItems(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {

	var mapItems []*models.MapItem
	rows, err := db.Queryx("SELECT * FROM mapItems")
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		mapItem := models.MapItem{}
		err = rows.Scan(&mapItem.Id,
			&mapItem.Name,
			&mapItem.Description,
			&mapItem.Image,
		)
		if err != nil {
			log.Fatalln(err)
		}
		mapItems = append(mapItems, &mapItem)
	}

	for _, mapItem := range mapItems {
		var location models.Location
		db.QueryRowx("SELECT id, lat, lng FROM locations WHERE locations.mapItem = ?", mapItem.Id).StructScan(&location)
		mapItem.Location = location

		var link models.Link
		rows, err := db.Queryx("SELECT id, name, type, link FROM links WHERE links.mapItemId = ?", mapItem.Id)
		if err != nil {
			log.Fatalln(err)
		}
		for rows.Next() {
			err = rows.StructScan(&link)
			if err != nil {
				fmt.Println(err)
			}
			mapItem.Links = append(mapItem.Links, link)
		}
	}

	out, err := json.Marshal(mapItems)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(out))
}

func experimental(db *sqlx.DB) {
}

func main() {

	db, err := sqlx.Open("sqlite3", "db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// test connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	experimental(db)

	http.HandleFunc("/api/getMapItems", MkRegisterHandler(db))
	//	http.HandleFunc("/api/getUserInfo", sendUserInfo)
	//	http.HandleFunc("/api/sendPosition", sendPrize)

	log.Fatal(http.ListenAndServe(":"+port, nil))

}

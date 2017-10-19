package mapItemsApi

import (
	"Loc_server/models"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

// MkGetMapItemsHandler returns handler for getMapItems request
func MkGetMapItemsHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getMapItems(w, r, db)
	}
}

func getMapItems(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	var mapItems []*models.MapItem
	rows, err := db.Queryx("SELECT * FROM mapItems")
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		mapItem := models.MapItem{}
		err = rows.StructScan(&mapItem)
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

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

func getMapItemsPartly(db *sqlx.DB) ([]*models.MapItem, error) {

	var mapItems []*models.MapItem

	rows, err := db.Queryx("SELECT * FROM mapItems")
	if err != nil {
		fmt.Println("Get mapItems from DB error: " + err.Error())
		return nil, err
	}

	for rows.Next() {
		mapItem := models.MapItem{}
		err = rows.StructScan(&mapItem)
		if err != nil {
			fmt.Println("MapItems struct error: " + err.Error())
			return nil, err
		}
		mapItems = append(mapItems, &mapItem)
	}
	return mapItems, nil
}

func getMapItemsLocations(mapItems []*models.MapItem, db *sqlx.DB) error {
	for _, mapItem := range mapItems {
		var location models.Location
		row := db.QueryRowx("SELECT id, lat, lng FROM locations WHERE locations.mapItem = ?", mapItem.Id)
		err := row.StructScan(&location)
		if err != nil {
			fmt.Println("Location struct error: " + err.Error())
			return err
		}
		mapItem.Location = location
	}
	return nil
}

func getMapItemsLinks(mapItems []*models.MapItem, db *sqlx.DB) error {
	for _, mapItem := range mapItems {
		var link models.Link
		rows, err := db.Queryx("SELECT id, name, type, link FROM links WHERE links.mapItemId = ?", mapItem.Id)
		if err != nil {
			fmt.Println("Get links from DB error: " + err.Error())
			return err
		}
		for rows.Next() {
			err = rows.StructScan(&link)
			if err != nil {
				fmt.Println("Links struct error: " + err.Error())
				return err
			}
			mapItem.Links = append(mapItem.Links, link)
		}
	}
	return nil
}

func getMapItems(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {

	var mapItems []*models.MapItem

	mapItems, err := getMapItemsPartly(db)
	if err != nil {
		log.Fatalln("Partly get mapItems error: " + err.Error())
	}

	err = getMapItemsLocations(mapItems, db)
	if err != nil {
		log.Fatalln("Get mapItems locations error: " + err.Error())
	}

	err = getMapItemsLinks(mapItems, db)
	if err != nil {
		log.Fatalln("Get mapItems links error: " + err.Error())
	}

	out, err := json.Marshal(mapItems)
	if err != nil {
		log.Fatalln("Marshal mapItems error: " + err.Error())
	}

	fmt.Fprintf(w, string(out))
}

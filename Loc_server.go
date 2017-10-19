package main

import (
	"Loc_server/achievementsApi"
	"Loc_server/mapItemsApi"
	"Loc_server/usersApi"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

func main() {

	db, err := sqlx.Open("sqlite3", "db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	// returns list of mapItem objects (TODO: return objects only from screen)
	http.HandleFunc("/api/getMapItems", mapItemsApi.MkGetMapItemsHandler(db))

	// returns info of user with requested id
	http.HandleFunc("/api/getUserInfo", usersApi.MkGetUserInfoHandler(db))

	// returns bonus object if user doesn't have this ahievement yet
	http.HandleFunc("/api/sendPosition", achievementsApi.MkSendBonusHandler(db))

	log.Fatal(http.ListenAndServe(":"+port, nil))

}

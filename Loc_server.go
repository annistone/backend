package main

import (
	"Loc_server/mapItemsApi"
	"Loc_server/usersApi"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

//sendUserInfo
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

	http.HandleFunc("/api/getMapItems", mapItemsApi.MkGetMapItemsHandler(db))
	http.HandleFunc("/api/getUserInfo", usersApi.MkGetUserInfoHandler(db))
	//	http.HandleFunc("/api/sendPosition", sendPrize)

	log.Fatal(http.ListenAndServe(":"+port, nil))

}

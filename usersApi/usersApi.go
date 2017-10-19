package usersApi

import (
	"Loc_server/models"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
)

// MkGetUserInfoHandler returns handler for getUserInfo request
func MkGetUserInfoHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getUserInfo(w, r, db)
	}
}

func getUserInfo(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	userIDStr := r.FormValue("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		fmt.Println("error:", err)
	}

	var user models.User
	db.QueryRowx("SELECT * FROM users WHERE users.id = ?", userID).StructScan(&user)

	var achievement models.Achievement
	rows, err := db.Queryx("SELECT achievements.id, text, description, image "+
		"FROM usersachievements "+
		"INNER JOIN achievements ON usersachievements.achievementId = achievements.Id "+
		"WHERE usersachievements.userId = ?", user.Id)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		err = rows.StructScan(&achievement)
		if err != nil {
			fmt.Println(err)
		}
		user.Achievements = append(user.Achievements, achievement)
	}

	out, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(out))
}

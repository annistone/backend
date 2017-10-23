package achievementsApi

import (
	"Loc_server/models"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"strconv"
)

// MkSendBonusHandler returns handler for sendPosition request
func MkSendBonusHandler(db *sqlx.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendBonus(w, r, db)
	}
}

//TODO: tokens
func sendBonus(w http.ResponseWriter, r *http.Request, db *sqlx.DB) {
	userIDStr := r.FormValue("userID")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Fatalln("error:", err)
	}

	achivementIDStr := r.FormValue("achievementID")
	receivedAchivementID, err := strconv.Atoi(achivementIDStr)
	if err != nil {
		log.Fatalln("error:", err)
	}

	var achievement models.Achievement

	rows, err := db.Queryx("SELECT achievements.id, text, description, image "+
		"FROM usersachievements "+
		"INNER JOIN achievements ON usersachievements.achievementId = achievements.Id "+
		"WHERE usersachievements.userId = ?", userID)
	if err != nil {
		log.Fatalln(err)
	}

	prices := make(chan int, 10)

	go func() {
		for rows.Next() {
			err = rows.StructScan(&achievement)
			if err != nil {
				log.Fatal(err)
			}
			if achievement.Id == receivedAchivementID {
				prices <- 0
			}
		}
		prices <- 100
	}()
	price := <-prices

	if price != 0 {
		rating := 0
		db.QueryRowx("SELECT rating FROM users WHERE users.id = ?", userID).Scan(&rating)
		rating += price
		tx := db.MustBegin()
		tx.MustExec("UPDATE users SET rating=? WHERE users.id = ?", rating, userID)
		err = tx.Commit()
		if err != nil {
			fmt.Println(err)
		}
		tx = db.MustBegin()
		tx.MustExec("INSERT INTO usersAchievements (userId, achievementId) VALUES (?, ?);", userID, receivedAchivementID)
		err = tx.Commit()
		if err != nil {
			fmt.Println(err)
		}
	}

	bonus := models.Bonus{
		Price:       price,
		Achievement: achievement,
	}

	out, err := json.Marshal(bonus)
	if err != nil {
		log.Fatal("error:", err)
	}

	fmt.Fprintf(w, string(out))
}

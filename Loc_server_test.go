package main_test

import (
	"Loc_server/achievementsApi"
	"Loc_server/mapItemsApi"
	"Loc_server/models"
	"Loc_server/usersApi"
	"encoding/json"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func mkTestGetUserInfo(db *sqlx.DB) func(t *testing.T) {

	return func(t *testing.T) {
		server := httptest.NewServer(usersApi.MkGetUserInfoHandler(db))
		reader := strings.NewReader("")

		request, err := http.NewRequest("GET", server.URL+"/api/getUserInfo?id=2", reader)

		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != 200 {
			t.Errorf("Success expected: %d", res.StatusCode)
		}

		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}

		var user models.User

		err = json.Unmarshal(bodyBytes, &user)

		if err != nil {
			t.Errorf("Unmarshal user error:" + err.Error())
		}

		t.Log(user)
	}
}

func mkTestGetMapItems(db *sqlx.DB) func(t *testing.T) {

	return func(t *testing.T) {
		server := httptest.NewServer(mapItemsApi.MkGetMapItemsHandler(db))
		reader := strings.NewReader("")

		request, err := http.NewRequest("GET", server.URL+"/api/getMapItems", reader)

		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != 200 {
			t.Errorf("Success expected: %d", res.StatusCode)
		}

		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}

		var mapItems []models.MapItem

		err = json.Unmarshal(bodyBytes, &mapItems)
		if err != nil {
			t.Errorf("Unmarshal mapItems error:" + err.Error())
		}

		t.Log(mapItems)
	}
}

func mkTestSendBonus(db *sqlx.DB) func(t *testing.T) {

	return func(t *testing.T) {
		tx := db.MustBegin()
		tx.MustExec("DELETE FROM usersAchievements WHERE userID = ? AND achievementID = ?", 1, 1)
		err := tx.Commit()
		if err != nil {
			t.Error(err)
		}

		server := httptest.NewServer(achievementsApi.MkSendBonusHandler(db))
		reader := strings.NewReader("")

		request, err := http.NewRequest("POST", server.URL+"/api/getMapItems?userID=1&achievementID=1", reader)

		res, err := http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != 200 {
			t.Errorf("Success expected: %d", res.StatusCode)
		}

		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}

		var bonus models.Bonus

		err = json.Unmarshal(bodyBytes, &bonus)
		if err != nil {
			t.Errorf("Unmarshal mapItems error:" + err.Error())
		}

		t.Log(bonus)

		if bonus.Price == 0 {
			t.Error("Logic error, want price != 0")
		}

		res, err = http.DefaultClient.Do(request)

		if err != nil {
			t.Error(err)
		}

		if res.StatusCode != 200 {
			t.Errorf("Success expected: %d", res.StatusCode)
		}

		bodyBytes, err = ioutil.ReadAll(res.Body)
		if err != nil {
			t.Error(err)
		}

		err = json.Unmarshal(bodyBytes, &bonus)
		if err != nil {
			t.Errorf("Unmarshal mapItems error:" + err.Error())
		}

		t.Log(bonus)

		if bonus.Price != 0 {
			t.Error("Logic error, want price = 0")
		}
	}
}

func TestServer(t *testing.T) {

	db, err := sqlx.Open("sqlite3", "db_test")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	t.Run("TestGetMapItems", mkTestGetMapItems(db))
	t.Run("TestGetUserInfo", mkTestGetUserInfo(db))
	t.Run("TestSendBonus", mkTestSendBonus(db))

}

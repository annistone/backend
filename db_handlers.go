package main

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"strconv"
)

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "tmpl/register.html")
		return
	}
	// grab user info
	username := r.FormValue("username")
	password := r.FormValue("password")
	role := r.FormValue("role")
	// Check existence of user
	var user User
	err := db.QueryRow("SELECT username, password, role FROM users WHERE username=?",
		username).Scan(&user.Username, &user.Password, &user.Role)
	switch {
	// user is available
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		checkInternalServerError(err, w)
		// insert to database
		_, err = db.Exec(`INSERT INTO users(username, password, role) VALUES(?, ?, ?)`,
			username, hashedPassword, role)
		fmt.Println("Created user: ", username)
		checkInternalServerError(err, w)
	case err != nil:
		http.Error(w, "loi: "+err.Error(), http.StatusBadRequest)
		return
	default:
		return
		//http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	}
}

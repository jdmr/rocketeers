package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

// User struct
type User struct {
	ID        string   `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Gender    string   `json:"gender"`
	Birthdate string   `json:"birthdate"`
	Email     string   `json:"email"`
	Phone     string   `json:"phone"`
	Carrier   string   `json:"carrier"`
	Roles     []string `json:"roles"`
}

func updateRegistration(c echo.Context) error {
	user := User{}
	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Error("Could not decode user: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not decode user: "+err.Error())
	}

	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	tx, err := conn.Begin()
	if err != nil {
		log.Error("Could not start transaction: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not start transaction: "+err.Error())
	}

	_, err = tx.Exec(`
		update users set first_name = ?, last_name = ?, gender = ?, birthdate = ?, phone = ?
		where email = ?
	`, user.FirstName, user.LastName, user.Gender, user.Birthdate, user.Phone, user.Email)
	if err != nil {
		log.Error("Could not update user: ", user.Email, " : ", err)
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, "Could not update user: "+user.Email+" : "+err.Error())
	}

	var id string
	err = tx.QueryRow(`
		select id from users where email = ?
	`, user.Email).Scan(&id)
	if err != nil {
		log.Error("Could not get user id for: ", user.Email, " : ", err)
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, "Could not get user id for: "+user.Email+" : "+err.Error())
	}

	for _, role := range user.Roles {
		_, err = tx.Exec(`
			insert into user_roles(user_id, role_id) values(?,?)
		`, id, role)
		if err != nil {
			log.Error("Could not add role to user: ", user.Email, " : ", id, " : ", role, " : ", err)
			tx.Rollback()
			return c.JSON(http.StatusInternalServerError, "Could not add role to user: "+user.Email+" : "+err.Error())
		}
	}

	tx.Commit()

	return c.NoContent(http.StatusOK)
}

package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func getRoles(c echo.Context) error {
	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	rows, err := conn.Query(`
		select id from roles
	`)
	if err != nil {
		log.Error("Could not get list of roles: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not get list of roles: "+err.Error())
	}

	roles := []string{}
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			log.Error("Could not get role: ", err)
			return c.JSON(http.StatusInternalServerError, "Could not get role: "+err.Error())
		}

		roles = append(roles, id)
	}

	return c.JSON(http.StatusOK, roles)
}

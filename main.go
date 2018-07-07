package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName(".rocketeers")
	viper.AddConfigPath("$HOME")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Could not read configuration file: ", err)
		panic(fmt.Errorf("Could not read configuration file: %s", err))
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	switch viper.GetString("log.level") {
	case "DEBUG":
		log.SetLevel(log.DEBUG)
	case "INFO":
		log.SetLevel(log.INFO)
	case "WARN":
		log.SetLevel(log.WARN)
	case "ERROR":
		log.SetLevel(log.ERROR)
	case "OFF":
		log.SetLevel(log.OFF)
	}

	log.Info("Rocketeers is up and running on port: 9000")
	e.Logger.Fatal(e.Start(":9000"))
}

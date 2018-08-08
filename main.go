package main

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	// MySQL DB
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	viper.SetConfigName(".rocketeers")
	viper.AddConfigPath("$HOME")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Could not read configuration file: ", err)
		panic(fmt.Errorf("Could not read configuration file: %s", err))
	}

	pem := viper.Get("key.public").(string)
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
	if err != nil {
		log.Error("Could not load pem", err)
		panic(err)
	}

	jwtConfig := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    key,
		SigningMethod: "RS256",
	})

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

	e.GET("/v1/auth", login)
	e.POST("/v1/auth", auth)

	rolesGroup := e.Group("/v1/roles")
	rolesGroup.Use(jwtConfig)
	rolesGroup.GET("", getRoles)

	registrationGroup := e.Group("/v1/registration")
	registrationGroup.Use(jwtConfig)
	registrationGroup.POST("", updateRegistration)

	e.POST("/v1/questions", addQuestionController)
	e.GET("/v1/questions", getQuestionsController)
	e.DELETE("/v1/questions/:questionID", deleteQuestionController)
	e.GET("/v1/questions/:questionID", getQuestionController)
	e.POST("/v1/questions/:questionID/answers", addAnswerController)
	e.DELETE("/v1/questions/:questionID/answers/:answerID", deleteAnswerController)
	e.POST("/v1/games", addGameController)
	e.GET("/v1/games", getGamesController)
	e.GET("/v1/games/:gameID", getGameController)
	e.POST("/v1/games/:gameID/teams", addTeamController)
	e.GET("/v1/games/:gameID/teams/:teamID", getTeamController)
	e.POST("/v1/games/:gameID/teams/:teamID/answers/:answerID", addTeamAnswerController)

	m := melody.New()
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		switch s.Request.URL.Path {
		case "/ws/pbe/games":
			games, err := getGames()
			if err != nil {
				log.Error("Could not get updated games: ", err)
				return
			}
			msg, _ = json.Marshal(games)
		}
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	}

	e.GET("/ws/pbe/games", func(c echo.Context) error {
		m.HandleRequest(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":9000"))
}

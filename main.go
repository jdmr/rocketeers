package main

import (
	"encoding/json"
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	melody "gopkg.in/olahol/melody.v1"
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

	// e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
	// 	Root:  "public",
	// 	HTML5: true,
	// }))

	// e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	if username == viper.GetString("auth.username") && password == viper.GetString("auth.password") {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))

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

	e.GET("/api/v1/auth", login)
	e.POST("/api/v1/auth", auth)

	rolesGroup := e.Group("/api/v1/roles")
	rolesGroup.Use(jwtConfig)
	rolesGroup.GET("", getRoles)

	registrationGroup := e.Group("/api/v1/registration")
	registrationGroup.Use(jwtConfig)
	registrationGroup.POST("", updateRegistration)

	e.POST("/api/v1/questions", addQuestionController)
	e.GET("/api/v1/questions", getQuestionsController)
	e.DELETE("/api/v1/questions/:questionID", deleteQuestionController)
	e.GET("/api/v1/questions/:questionID", getQuestionController)
	e.POST("/api/v1/questions/:questionID/answers", addAnswerController)
	e.DELETE("/api/v1/questions/:questionID/answers/:answerID", deleteAnswerController)
	e.POST("/api/v1/games", addGameController)
	e.GET("/api/v1/games", getGamesController)
	e.DELETE("/api/v1/games/:gameID", deleteGameController)
	e.GET("/api/v1/games/:gameID", getGameController)
	e.POST("/api/v1/games/:gameID/teams", addTeamController)
	e.GET("/api/v1/games/:gameID/teams/:teamID", getTeamController)
	e.DELETE("/api/v1/games/:gameID/teams/:teamID/answers/:answerID", deleteTeamAnswerController)
	e.POST("/api/v1/games/:gameID/teams/:teamID/answers/:answerID", addTeamAnswerController)
	e.POST("/api/v1/games/:gameID/start", startGameController)
	e.POST("/api/v1/games/:gameID/finish", finishGameController)
	e.GET("/api/v1/games/:gameID/finished", getFinishedGameController)
	e.POST("/api/v1/games/:gameID/next", nextQuestionController)
	e.POST("/api/v1/games/:gameID/previous", previousQuestionController)
	e.GET("/api/v1/games/:gameID/current", getCurrentQuestionController)
	e.GET("/api/v1/games/:gameID/home", getHomeTeamController)

	m := melody.New()
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		if strings.HasPrefix(s.Request.URL.Path, "/ws/pbe/teams") {
			games, err := getGames()
			if err != nil {
				log.Error("Could not get updated games: ", err)
				return
			}
			msg, _ = json.Marshal(games)
		} else if strings.HasPrefix(s.Request.URL.Path, "/ws/pbe/game") {
			gameID := string(msg)
			log.Info("WS: getting current question for : ", gameID)
			question, err := getCurrentQuestion(gameID)
			if err != nil {
				log.Error("Could not get the current question: ", err)
				return
			}
			msg, _ = json.Marshal(question)
		}
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

	e.GET("/ws/pbe/teams", func(c echo.Context) error {
		m.HandleRequest(c.Response(), c.Request())
		return nil
	})

	e.GET("/ws/pbe/game/:gameID", func(c echo.Context) error {
		keys := make(map[string]interface{})
		keys["gameID"] = c.Param("gameID")
		m.HandleRequestWithKeys(c.Response(), c.Request(), keys)
		return nil
	})

	e.Logger.Fatal(e.Start(":9000"))
}

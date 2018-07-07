package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func login(c echo.Context) error {
	state := c.Param("state")
	log.Info("Logging in with Google: ", state)

	b, err := ioutil.ReadFile("config/google.json")
	if err != nil {
		log.Error("Couldn't read google's configuration file: ", err)
	}
	googleConfig, err := google.ConfigFromJSON(b, "email")
	if err != nil {
		log.Error("Couldn't parse google's configuration file: ", err)
	}
	googleConfig.Scopes = []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
	}
	googleConfig.RedirectURL = "http://localhost:8080/rocketeers/login"

	url := googleConfig.AuthCodeURL(state)
	log.Info("Redirecting to: ", url)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func auth(c echo.Context) error {
	state := c.Param("state")
	code, err := url.QueryUnescape(c.Param("code"))
	if err != nil {
		log.Error("Couldn't read code : ", c.Param("code"), " : ", err)
	}
	log.Info("Code: ", code)

	b, err := ioutil.ReadFile("config/google.json")
	if err != nil {
		log.Error("Couldn't read google's configuration file: ", err)
	}
	googleConfig, err := google.ConfigFromJSON(b, "email")
	if err != nil {
		log.Error("Couldn't parse google's configuration file: ", err)
	}
	googleConfig.Scopes = []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email",
	}
	googleConfig.RedirectURL = "http://localhost:8080/rocketeers/login"

	t, err := googleConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Error("Could not exchange token: ", err)
		message := url.QueryEscape("Could not exchange token")
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8080/rocketeers/login?error="+message)
	}

	resp, err := http.Get("https://www.googleapis.com/plus/v1/people/me?access_token=" + t.AccessToken)
	if err != nil {
		log.Error("Could not get the user's information: ", err)
		message := url.QueryEscape("Could not get the user's information")
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8080/rocketeers/login?error="+message)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Could not read the user's information: ", err)
		message := url.QueryEscape("Could not read the user's information")
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8080/rocketeers/login?error="+message)
	}
	var googleUser GoogleUser
	err = json.Unmarshal(body, &googleUser)
	if err != nil {
		log.Error("Could not parse the user's information: ", err)
		message := url.QueryEscape("Could not parse the user's information")
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:8080/rocketeers/login?error="+message)
	}

	url := "http://localhost:8080/login/google?state=" + state + "&token=token"
	log.Info("Redirecting to: ", url)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

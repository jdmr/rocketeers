package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// TokenFilter struct
type TokenFilter struct {
	Code string `json:"code"`
	Dev  bool   `json:"dev"`
}

// GoogleUserEmails struct
type GoogleUserEmails struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

// GoogleUserName struct
type GoogleUserName struct {
	FamilyName string `json:"familyName"`
	GivenName  string `json:"givenName"`
}

// GoogleUserImage struct
type GoogleUserImage struct {
	URL string `json:"url"`
}

// GoogleUser struct
type GoogleUser struct {
	Name   *GoogleUserName     `json:"name"`
	Emails []*GoogleUserEmails `json:"emails"`
	Image  *GoogleUserImage    `json:"image"`
	Gender string              `json:"gender"`
}

func login(c echo.Context) error {
	state, _ := UUID()
	dev := c.FormValue("dev")
	b, err := ioutil.ReadFile("config/google.json")
	if err != nil {
		log.Error("Could not read google config file: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not read google config file: "+err.Error())
	}
	googleConfig, err := google.ConfigFromJSON(b, "email")
	if err != nil {
		log.Error("Could not parse google config file: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not parse google config file: "+err.Error())
	}
	googleConfig.Scopes = []string{
		"profile",
		"email",
	}
	if len(dev) > 0 {
		googleConfig.RedirectURL = "http://localhost:8080/rocketeers/login"
	} else {
		googleConfig.RedirectURL = "https://keenechurch.org/rocketeers/login"
	}

	url := googleConfig.AuthCodeURL(state)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func auth(c echo.Context) error {
	filter := TokenFilter{}
	err := json.NewDecoder(c.Request().Body).Decode(&filter)
	if err != nil {
		log.Error("Could not get token parameters: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not get token parameters: "+err.Error())
	}
	var redirectURI string
	if filter.Dev {
		redirectURI = "http://localhost:8080/rocketeers/login"
	} else {
		redirectURI = "https://keenechurch.org/rocketeers/login"
	}
	b, err := ioutil.ReadFile("config/google.json")
	if err != nil {
		log.Error("Could not read google config file: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not read google config file: "+err.Error())
	}
	googleConfig, err := google.ConfigFromJSON(b, "email")
	if err != nil {
		log.Error("Could not parse google config file: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not parse google config file: "+err.Error())
	}
	googleConfig.Scopes = []string{
		"profile",
		"email",
	}
	googleConfig.RedirectURL = redirectURI

	t, err := googleConfig.Exchange(oauth2.NoContext, filter.Code)
	if err != nil {
		log.Error("Could not exchange token for code: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not exchange token for code: "+err.Error())
	}

	resp, err := http.Get("https://www.googleapis.com/plus/v1/people/me?access_token=" + t.AccessToken)
	if err != nil {
		log.Error("Could not get user's information: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not get user's information: "+err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error("Could not read user's information: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not read user's information: "+err.Error())
	}
	var googleUser GoogleUser
	err = json.Unmarshal(body, &googleUser)
	if err != nil {
		log.Error("Could not parse user's information: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not parse user's information: "+err.Error())
	}

	log.Info("Google User", googleUser)
	email := googleUser.Emails[0].Value

	scopes := []string{}
	scopes = append(scopes, "openid")

	var (
		id        string
		firstName string
		lastName  string
		gender    string
		image     string
	)

	conn, err := sql.Open("mysql", viper.GetString("database.url"))
	if err != nil {
		log.Error("Open connection failed: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not open database : "+err.Error())
	}
	defer conn.Close()

	err = conn.QueryRow(`
		select
			id
			, first_name
			, last_name
			, coalesce(gender, 'male')
			, coalesce(image_url, '')
		from users
		where email = ?
	`, email).Scan(&id, &firstName, &lastName, &gender, &image)
	if err != nil {
		log.Info("User: ", email, " not found... Creating...", err)
		id, _ = UUID()
		firstName = googleUser.Name.GivenName
		lastName = googleUser.Name.FamilyName
		gender = googleUser.Gender
		image = googleUser.Image.URL

		_, err = conn.Exec(`
			insert into users(id, first_name, last_name, email, gender, image_url)
			values(?,?,?,?,?,?)
		`, id, firstName, lastName, email, gender, image)
		if err != nil {
			log.Error("Could not create user: ", email, " : ", err)
			return c.JSON(http.StatusInternalServerError, "Could not create user: "+email+" : "+err.Error())
		}
	}

	rows, err := conn.Query(`
		select 
			role_id
		from user_roles
		where user_id = ?
	`, id)
	if err != nil {
		log.Error("Could not get user roles: ", err)
		return c.JSON(http.StatusInternalServerError, "Could not get user roles: "+err.Error())
	}

	for rows.Next() {
		var scope string
		err = rows.Scan(&scope)
		if err != nil {
			log.Error(http.StatusInternalServerError, "Could not read user scope: ", err)
			return c.JSON(http.StatusInternalServerError, "Could not read user scope: "+err.Error())
		}
		scopes = append(scopes, scope)
	}

	token := jwt.New(jwt.SigningMethodRS256)

	clientID := "rocketeers"
	grantType := "authorization_code"

	log.Info("Host: ", c.Request().Host)
	issuedTime := time.Now()
	issued := issuedTime.Unix()
	expires := issuedTime.Add(time.Hour * 168).Unix()
	claims := token.Claims.(jwt.MapClaims)
	claims["scope"] = scopes
	claims["client_id"] = clientID
	claims["cid"] = clientID
	claims["azp"] = clientID
	claims["grant_type"] = grantType
	claims["user_id"] = email
	claims["image_url"] = image
	claims["name"] = firstName + " " + lastName
	claims["first_name"] = firstName
	claims["last_name"] = lastName
	claims["username"] = email
	claims["user_name"] = email
	claims["email"] = email
	claims["gender"] = gender
	claims["auth_time"] = issued
	claims["iat"] = issued
	claims["exp"] = expires
	claims["iss"] = c.Scheme() + "://" + c.Request().Host + c.Path()
	claims["aud"] = []string{"openid", clientID}

	pem := viper.Get("key.private").(string)
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(pem))
	if err != nil {
		log.Error("Could not load pem: ", err)
		return err
	}
	tk, err := token.SignedString(key)
	if err != nil {
		log.Error("Could not sign key: ", err)
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"access_token": tk,
		"token_type":   "bearer",
		"expires_in":   strconv.FormatInt(expires, 10),
	})
}

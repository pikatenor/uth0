package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Env struct {
	Port   int    `default:"8080"`
	User   string `required:"true"`
	Pass   string `required:"true"`
	Secret string `required:"true"`
}

type HasuraCustomClaims struct {
	Hasura HasuraCustomFields `json:"https://hasura.io/jwt/claims"`
	jwt.StandardClaims
}

type HasuraCustomFields struct {
	AllowedRoles []string `json:"x-hasura-allowed-roles"`
	DefaultRole  string   `json:"x-hasura-default-role"`
}

func main() {
	var v Env
	err := envconfig.Process("", &v)
	if err != nil {
		log.Fatal(err.Error())
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/auth", auth)
	e.GET("/", health)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", v.Port)))
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "check")
}

func auth(c echo.Context) error {
	var v Env
	envconfig.Process("", &v)

	user := c.FormValue("user")
	pass := c.FormValue("pass")

	if user != v.User || pass != v.Pass {
		return echo.ErrUnauthorized
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, &HasuraCustomClaims{
		HasuraCustomFields{
			AllowedRoles: []string{"user"},
			DefaultRole:  "user",
		},
		jwt.StandardClaims{
			Issuer:    "uso/0.1",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
		},
	})

	token, err := tokenClaim.SignedString([]byte(v.Secret))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}

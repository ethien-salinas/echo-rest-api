package handler

import (
	"github.com/ethien-salinas/echo-rest-api/db"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// Login validate credentials to get jwt token
func Login(c echo.Context) error {

	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	// Throws unauthorized error
	userLogged := db.ValidateUser(u.Email, u.Password)

	if userLogged == "" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = userLogged
	claims["admin"] = true
	claims["exp"] = time.Now().Add(3 * time.Hour).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

package handler

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// SaveUser - POST method
func SaveUser(c echo.Context) error {
	return c.JSON(http.StatusCreated, User{Name: "Ethien", Email: "ethien@intelligential.tech"})
}

// GetUser - GET method
func GetUser(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	log.Println("[0] *** " + user.Raw)
	claims := user.Claims.(jwt.MapClaims)

	keys := make([]string, 0, len(claims))
	for k := range claims {
		keys = append(keys, k)
	}
	log.Print("[1] *** ")
	log.Println(keys)

	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

// UpdateUser - PUT method
func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusProcessing, id)
}

// DeleteUser - DELETE method
func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.JSON(http.StatusAccepted, id)
}

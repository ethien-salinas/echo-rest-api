package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ethien-salinas/echo-rest-api/handler"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Echo instance
	e := echo.New()
	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routing
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Echo.v4 REST API")
	})
	e.POST("/login", handler.Login)
	// Restricted group
	r := e.Group("/users")
	r.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	r.POST("/", handler.SaveUser)
	r.GET("/:id", handler.GetUser)
	r.PUT("/:id", handler.UpdateUser)
	r.DELETE("/:id", handler.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}

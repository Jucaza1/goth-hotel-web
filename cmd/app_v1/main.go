package main

import (
	"io"
	"log"
	"net/http"

	"github.com/Jucaza1/goth-hotel-web/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	//const server = "http://localhost:4000"
	const server = "http://localhost:3000/echo"

	app := echo.New()
	app.HTTPErrorHandler = handler.ErrorHandler
	userHandler := handler.NewUserHandler(server)
	homeHandler := handler.NewHomeHandler(server)
	loginHandler := handler.NewLoginHandler(server)

	app.Any("/public/*", public())
	app.GET("/", homeHandler.HandleHomeShow)
	app.GET("/login", loginHandler.HandleLoginShow)
	app.GET("/user", userHandler.HandleUserShow)
	app.Any("/echo/*", func(c echo.Context) error {
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to read request body")
		}

		// Create a response with the same body
		response := string(body)

		// Mirror the request headers
		for name, values := range c.Request().Header {
			for _, value := range values {
				c.Response().Header().Add(name, value)
			}
		}

		// Set the response content type to the same as the request content type
		c.Response().Header().Set(echo.HeaderContentType, c.Request().Header.Get(echo.HeaderContentType))

		return c.String(http.StatusBadRequest, response)
	})

	log.Panic(app.Start(":3000"))

}

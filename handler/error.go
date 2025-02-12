package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("public/HTML/%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
        return
	}
	c.Logger().Error(err)
}

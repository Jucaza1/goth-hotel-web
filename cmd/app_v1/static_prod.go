//go:build !dev
// +build !dev

package main

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed public
var publicFS embed.FS

func public() echo.HandlerFunc {
	return echo.WrapHandler(http.StripPrefix("/static/", http.FileServerFS(publicFS)))
}

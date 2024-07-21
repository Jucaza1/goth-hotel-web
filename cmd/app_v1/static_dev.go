//go:build dev
// +build dev

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func public() echo.HandlerFunc {
	fmt.Println("building static files for development")
	//return echo.StaticDirectoryHandler(os.DirFS("public"), false)
	return echo.WrapHandler(http.StripPrefix("/public", http.FileServerFS(os.DirFS("public"))))
}

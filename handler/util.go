package handler

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(c echo.Context, status int, component templ.Component) error {
	c.Response().Status = status
	return component.Render(c.Request().Context(), c.Response())
}

// type HTTPHandler func(http.ResponseWriter, *http.Request) error

// func Wrap(h HTTPHandler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if err := h(w, r); err != nil {
// 			slog.Error("HTTP handler error", "err", err, "path", r.URL.Path)
// 		}
// 	}
// }

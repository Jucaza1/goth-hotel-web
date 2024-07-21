package handler

import (
	"github.com/Jucaza1/goth-hotel-web/view/login"
	"github.com/labstack/echo/v4"
)

type LoginHandler struct {
	server string
}

func NewLoginHandler(server string) *LoginHandler {
	return &LoginHandler{
		server: server,
	}
}

func (h *LoginHandler) HandleLoginShow(c echo.Context) error {
	return render(c, 200, login.Show(h.server))
}

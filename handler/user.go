package handler

import (
	"github.com/Jucaza1/goth-hotel-web/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	server string
}

func NewUserHandler(server string) UserHandler {
	return UserHandler{
		server: server,
	}
}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	return render(c, 200, user.Show())
}

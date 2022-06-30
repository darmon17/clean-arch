package routes

import (
	"project2/factory"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()
	e.GET("/users", presenter.UserPrsenter.GetAll)
	return e
}

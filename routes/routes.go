package routes

import (
	"project2/factory"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {
	e := echo.New()
	e.GET("/users", presenter.UserPrsenter.GetAll)
	e.POST("/login", presenter.UserPrsenter.LoginAuth)
	e.POST("/users", presenter.UserPrsenter.PostUser)
	return e
}

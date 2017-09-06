package router

import (
	"github.com/float1251/echo_sample/controller"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
)

func SetRouting(e *echo.Echo, db *gorm.DB) {
	u := controller.NewUserHandler(db)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.GET("/not_restricted/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World Open")
	})
	e.POST("/user/login/", u.Login)
	e.POST("/user/create/", u.UserCreate)
}

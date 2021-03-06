package router

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/float1251/echo_sample/controller"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"net/http"
)

func SetRouting(e *echo.Echo, db *gorm.DB) {
	u := controller.NewUserHandler(db)
	e.GET("/", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		id := claims["id"].(string)
		return c.String(http.StatusOK, "Welcome "+id+"!")
	})
	e.GET("/not_restricted/", func(c echo.Context) error {
		b := controller.BaseResponse{}
		b.Code = 200
		b.Message = controller.UserLoginResponse{}
		return c.JSON(http.StatusOK, b)
	})
	e.POST("/user/login/", u.Login)
	e.POST("/user/create/", u.UserCreate)
}

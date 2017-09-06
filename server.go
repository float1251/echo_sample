package main

import (
	"github.com/float1251/echo_sample/model"
	"github.com/float1251/echo_sample/router"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/user/create/" ||
				c.Path() == "/user/login/" ||
				c.Path() == "/not_restricted/" {
				return true
			}
			return false
		},
	}))
	e.Use(middleware.Recover())

	// db
	db, err := gorm.Open("sqlite3", "/tmp/gorm.db?cache=shared")
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(10)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// migration
	model.Migrate(db)

	router.SetRouting(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}

package controller

import (
	"github.com/float1251/echo_sample/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"strconv"
)

type handler struct {
	db *gorm.DB
}

type UserHandler struct {
	handler
}

type UserCreateResponse struct {
	ID   uint
	Uuid uuid.UUID
	Name string
}

type UserLoginResponse struct {
}

func NewUserHandler(d *gorm.DB) *UserHandler {
	u := new(UserHandler)
	u.db = d
	return u
}

func (h *UserHandler) Login(c echo.Context) error {
	u := &model.UserModel{}
	h.db.Where(&model.UserModel{Name: c.Param("id")}).First(u)
	return c.String(http.StatusOK, "User: "+strconv.FormatUint(uint64(u.ID), 10)+", Name: "+u.Name)
}

func (h *UserHandler) UserCreate(c echo.Context) error {
	id := c.Param("id")
	u := model.NewUserModel(id)
	h.db.Create(u)
	res := &UserCreateResponse{ID: u.ID, Uuid: u.Uuid, Name: u.Name}
	return c.JSON(http.StatusOK, res)
}

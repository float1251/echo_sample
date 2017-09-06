package controller

import (
	"encoding/json"
	"fmt"
	"github.com/float1251/echo_sample/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
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

type UserCreateRequest struct {
	Password string `json:"password"`
	UserName string `json:"user_name"`
}

type UserLoginResponse struct {
	Token string
}

type UserLoginRequest struct {
	ID       string
	Password string
}

func NewUserHandler(d *gorm.DB) *UserHandler {
	u := new(UserHandler)
	u.db = d
	return u
}

func (h *UserHandler) Login(c echo.Context) error {
	req := new(UserLoginRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	u := &model.UserModel{}
	h.db.Where(&model.UserModel{Name: req.ID}).First(u)
	res := UserLoginResponse{Token: string(u.Password)}
	return c.JSON(http.StatusOK, res)
}

func (h *UserHandler) UserCreate(c echo.Context) error {
	req := new(UserCreateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	Log(req)
	u := model.NewUserModel(req.UserName, []byte(req.Password))
	h.db.Create(u)
	res := &UserCreateResponse{ID: u.ID, Uuid: u.Uuid, Name: u.Name}
	return c.JSON(http.StatusOK, res)
}

func Log(i interface{}) {
	j, _ := json.Marshal(i)
	fmt.Println(string(j))
}

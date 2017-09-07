package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/float1251/echo_sample/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"time"
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
	Token string `json:"token"`
}

type UserLoginRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
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

	if err := login(h, req.ID, req.Password); err != nil {
		return err
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = req.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Generate encoded token and send it as response.
	// TODO: change secret on production
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	res := UserLoginResponse{Token: t}
	return c.JSON(http.StatusOK, res)
}

func login(h *UserHandler, id string, password string) error {
	u := new(model.UserModel)
	id_uint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return errors.New("Parse Failed")
	}
	u.ID = uint(id_uint)
	res := new(model.UserModel)
	h.db.Where(u).First(res)
	if res.ID == 0 {
		return errors.New("Not Exist")
	}
	if err = bcrypt.CompareHashAndPassword(res.Password, []byte(password)); err != nil {
		return err
	}

	return nil
}

func (h *UserHandler) UserCreate(c echo.Context) error {
	req := new(UserCreateRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	// Log(req)
	u := model.NewUserModel(req.UserName, []byte(req.Password))
	h.db.Create(u)
	res := &UserCreateResponse{ID: u.ID, Uuid: u.Uuid, Name: u.Name}
	return c.JSON(http.StatusOK, res)
}

func Log(i interface{}) {
	j, _ := json.Marshal(i)
	fmt.Println(string(j))
}

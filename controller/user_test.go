package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"
	"encoding/json"
	"github.com/float1251/echo_sample/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"os"
)

var (
	db *gorm.DB
	h  *UserHandler
)

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	db, _ = gorm.Open("sqlite3", "/tmp/gorm_test.db")
	model.Migrate(db)
	h = NewUserHandler(db)
	res := m.Run()
	db.Close()
	os.Remove("/tmp/gorm_test.db")
	os.Exit(res)
}

func TestUserCreate(t *testing.T) {
	e := echo.New()
	arg := &UserCreateRequest{Password: "test", Name: "ttt"}

	body, _ := json.Marshal(arg)
	req := httptest.NewRequest(echo.POST, "/user/create/124", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.UserCreate(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		u := &model.UserModel{}
		tmp := &model.UserModel{}
		tmp.ID = 1
		db.Where(tmp).First(u)
		assert.Equal(t, uint(1), u.Model.ID)
		assert.Equal(t, "ttt", u.Name)
	}
}

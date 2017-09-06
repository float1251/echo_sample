package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	gorm.Model
	Uuid       uuid.UUID `gorm:"index"`
	Name       string
	RedDiamond int64
	Password   []byte
}

func NewUserModel(name string, password []byte) *UserModel {
	u := &UserModel{Uuid: uuid.NewV1(), Name: name}
	pass, _ := bcrypt.GenerateFromPassword(password, 5)
	u.Password = pass
	return u
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&UserModel{})
}

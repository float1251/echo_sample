package model

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type UserModel struct {
	gorm.Model
	Uuid       uuid.UUID `gorm:"index"`
	Name       string
	RedDiamond int64
}

func NewUserModel(name string) *UserModel {
	u := &UserModel{Uuid: uuid.NewV1(), Name: name}
	return u
}

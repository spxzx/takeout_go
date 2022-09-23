package model

import (
	"TakeOut/router/middleware"
	"gorm.io/gorm"
	"reflect"
)

type User struct {
	ID       int    `json:"id,string"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	IdNumber string `json:"idNumber"`
	Avatar   int    `json:"avatar"`
	Status   int    `json:"status"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Universal
}

func (u *User) IsEmpty() bool {
	return reflect.DeepEqual(u, &User{})
}

func (u *User) BeforeCreate(*gorm.DB) (err error) {
	u.CreateUser = middleware.CurrentId
	u.UpdateUser = middleware.CurrentId
	return
}

func (u *User) BeforeUpdate(*gorm.DB) (err error) {
	u.UpdateUser = middleware.CurrentId
	return
}

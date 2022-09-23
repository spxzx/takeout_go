package model

import (
	"TakeOut/router/middleware"
	"gorm.io/gorm"
	"reflect"
	"time"
)

type ShoppingCart struct {
	ID         int       `json:"id,string"`
	Name       string    `json:"name"`
	Image      string    `json:"image"`
	UserId     int       `json:"userId,string"`
	DishId     int       `json:"dishId,string"`
	SetmealId  int       `json:"setmealId,string"`
	DishFlavor string    `json:"dishFlavor"`
	Number     int       `json:"number"`
	Amount     float64   `json:"amount"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`
}

type ShoppingCartList []ShoppingCart

func (s *ShoppingCart) IsEmpty() bool {
	return reflect.DeepEqual(s, &ShoppingCart{})
}

func (s *ShoppingCart) BeforeCreate(*gorm.DB) (err error) {
	s.UserId = middleware.CurrentId
	return
}

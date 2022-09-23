package model

import (
	"TakeOut/router/middleware"
	"gorm.io/gorm"
	"reflect"
)

type SetmealDish struct {
	ID        int     `json:"id,string"`
	SetmealId int     `json:"setmealId"`
	DishId    int     `json:"dishId,string"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Copies    int     `json:"copies"`
	Sort      int     `json:"sort"`
	Universal
}

type SetmealDishList []SetmealDish

func (s *SetmealDish) IsEmpty() bool {
	return reflect.DeepEqual(s, &SetmealDish{})
}

func (s *SetmealDish) BeforeCreate(*gorm.DB) (err error) {
	s.SetmealId = CurSetmealId
	s.CreateUser = middleware.CurrentId
	s.UpdateUser = middleware.CurrentId
	return
}

func (s *SetmealDish) BeforeUpdate(*gorm.DB) (err error) {
	s.UpdateUser = middleware.CurrentId
	return
}

type SetmealDishDto struct {
	SetmealDish
	Image       string `json:"image"`
	Description string `json:"description"`
}

type SetmealDishDtoList []SetmealDishDto

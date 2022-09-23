package model

import (
	"TakeOut/router/middleware"
	"gorm.io/gorm"
	"reflect"
	"time"
)

type DishFlavor struct {
	ID         int       `json:"id,string"`
	DishId     int       `json:"dishId,string"`
	Name       string    `json:"name"`
	Value      string    `json:"value"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"updateTime"`
	CreateUser int       `json:"createUser"`
	UpdateUser int       `json:"updateUser"`
}

type DishFlavorList []DishFlavor

func (d *DishFlavor) IsEmpty() bool {
	return reflect.DeepEqual(d, &DishFlavor{})
}

func (d *DishFlavor) BeforeCreate(*gorm.DB) (err error) {
	d.DishId = CurDishId
	d.CreateUser = middleware.CurrentId
	d.UpdateUser = middleware.CurrentId
	return
}

func (d *DishFlavor) BeforeUpdate(*gorm.DB) (err error) {
	d.UpdateUser = middleware.CurrentId
	return
}

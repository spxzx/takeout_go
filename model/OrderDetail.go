package model

import (
	"gorm.io/gorm"
	"reflect"
)

type OrderDetail struct {
	ID         int     `json:"id,string"`
	Name       string  `json:"name"`
	Image      string  `json:"image"`
	OrderId    int     `json:"orderId,string"`
	DishId     int     `json:"dishId,string"`
	SetmealId  int     `json:"setmealId,string"`
	DishFlavor string  `json:"dishFlavor"`
	Number     int     `json:"number"`
	Amount     float64 `json:"amount"`
}

type OrderDetailList []OrderDetail

func (o *OrderDetail) IsEmpty() bool {
	return reflect.DeepEqual(o, &OrderDetail{})
}

func (o *OrderDetail) BeforeCreate(*gorm.DB) (err error) {
	o.OrderId = CurOrderId
	return
}

package model

import (
	"gorm.io/gorm"
	"reflect"
	"time"
)

type Orders struct {
	ID            int       `json:"id,string"`
	Number        string    `json:"number"`
	Status        int       `json:"status"`
	UserId        int       `json:"userId,string"`
	AddressBookId int       `json:"addressBookId,string"`
	OrderTime     time.Time `gorm:"autoCreateTime" json:"orderTime"`
	CheckoutTime  time.Time `gorm:"autoCreateTime" json:"checkoutTime"`
	PayMethod     int       `json:"payMethod"`
	Amount        float64   `json:"amount"`
	Remark        string    `json:"remark"`
	Phone         string    `json:"phone"`
	Address       string    `json:"address"`
	UserName      string    `json:"userName"`
	Consignee     string    `json:"consignee"`
}

var CurOrderId int

type OrdersList []Orders

func (o *Orders) IsEmpty() bool {
	return reflect.DeepEqual(o, &Orders{})
}

func (o *Orders) AfterCreate(*gorm.DB) (err error) {
	CurOrderId = o.ID
	return
}

func (o *Orders) AfterUpdate(*gorm.DB) (err error) {
	CurOrderId = o.ID
	return
}

type OrderDto struct {
	Orders
	OrderDetails OrderDetailList `json:"orderDetails"`
	Username     string          `json:"username"`
	Address      string          `json:"address"`
	SunNum       int             `json:"sunNum"`
}

type OrderDtoList []OrderDto

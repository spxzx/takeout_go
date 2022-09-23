package model

import (
	"TakeOut/router/middleware"
	"gorm.io/gorm"
	"reflect"
)

type AddressBook struct {
	ID           int    `json:"id,string"`
	UserId       int    `json:"userId,string"`
	Consignee    string `json:"consignee"`
	Sex          int    `json:"sex,string"`
	Phone        string `json:"phone"`
	ProvinceCode string `json:"provinceCode"`
	ProvinceName string `json:"provinceName"`
	CityCode     string `json:"cityCode"`
	CityName     string `json:"cityName"`
	DistrictCode string `json:"districtCode"`
	DistrictName string `json:"districtName"`
	Detail       string `json:"detail"`
	Label        string `json:"label"`
	IsDefault    int    `json:"isDefault"`
	Universal
}

type AddressBookList []AddressBook

func (a *AddressBook) IsEmpty() bool {
	return reflect.DeepEqual(a, &User{})
}

func (a *AddressBook) BeforeCreate(*gorm.DB) (err error) {
	a.UserId = middleware.CurrentId
	a.CreateUser = middleware.CurrentId
	a.UpdateUser = middleware.CurrentId
	return
}

func (a *AddressBook) BeforeUpdate(*gorm.DB) (err error) {
	a.UpdateUser = middleware.CurrentId
	return
}

package model

import (
	"TakeOut/router/middleware"
	"gorm.io/gorm"
	"reflect"
)

type Dish struct {
	ID          int     `json:"id,string"`
	Name        string  `json:"name"`
	CategoryId  int     `json:"categoryId,string"`
	Price       float64 `json:"price"`
	Code        string  `json:"code"`
	Image       string  `json:"image"`
	Description string  `json:"description"`
	Status      int     `json:"status"`
	Sort        int     `json:"sort"`
	Universal
}

var CurDishId int

type DishList []Dish

func (d *Dish) IsEmpty() bool {
	return reflect.DeepEqual(d, &Dish{})
}

func (d *Dish) BeforeCreate(*gorm.DB) (err error) {
	d.CreateUser = middleware.CurrentId
	d.UpdateUser = middleware.CurrentId
	return
}

func (d *Dish) AfterCreate(*gorm.DB) (err error) {
	CurDishId = d.ID
	return
}

func (d *Dish) BeforeUpdate(*gorm.DB) (err error) {
	d.UpdateUser = middleware.CurrentId
	return
}

func (d *Dish) AfterUpdate(*gorm.DB) (err error) {
	CurDishId = d.ID
	return
}

type DishDto struct {
	Dish
	Flavors      DishFlavorList `json:"flavors"`
	CategoryName string         `json:"categoryName"`
	Copies       int            `json:"copies"`
}

type DishDtoList []DishDto

func GetDishWithFlavors(id int) DishDto {
	var dto DishDto
	GetOne(&dto.Dish, []string{"id"}, id)
	List(&dto.Flavors, "", []string{"dish_id"}, id)
	return dto
}

func UpdateDishWithFlavors(dto DishDto) {
	Update(dto.Dish)
	DeleteBatch(DishFlavor{}, "dish_id", dto.ID)
	InsertBatch(dto.Flavors)
}

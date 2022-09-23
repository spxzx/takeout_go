package model

import (
	"TakeOut/router/middleware"
	"gorm.io/gorm"
	"reflect"
)

type Setmeal struct {
	ID          int     `json:"id,string"`
	CategoryId  int     `json:"categoryId,string"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Status      int     `json:"status"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Universal
}

var CurSetmealId int

type SetmealList []Setmeal

func (s *Setmeal) IsEmpty() bool {
	return reflect.DeepEqual(s, &Setmeal{})
}

func (s *Setmeal) BeforeCreate(*gorm.DB) (err error) {
	s.CreateUser = middleware.CurrentId
	s.UpdateUser = middleware.CurrentId
	return
}

func (s *Setmeal) AfterCreate(*gorm.DB) (err error) {
	CurSetmealId = s.ID
	return
}

func (s *Setmeal) BeforeUpdate(*gorm.DB) (err error) {
	s.UpdateUser = middleware.CurrentId
	return
}

func (s *Setmeal) AfterUpdate(*gorm.DB) (err error) {
	CurSetmealId = s.ID
	return
}

type SetmealDto struct {
	Setmeal
	SetmealDishes SetmealDishList `json:"setmealDishes"`
	CategoryName  string          `json:"categoryName"`
}

type SetmealDtoList []SetmealDto

func GetSetmealWithDishes(id int) SetmealDto {
	var dto SetmealDto
	GetOne(&dto.Setmeal, []string{"id"}, id)
	List(&dto.SetmealDishes, "", []string{"setmeal_id"}, id)
	return dto
}

func UpdateSetmealWithDishes(dto SetmealDto) {
	Update(dto.Setmeal)
	DeleteBatch(SetmealDish{}, "setmeal_id", dto.ID)
	InsertBatch(dto.SetmealDishes)
}

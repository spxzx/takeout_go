package model

import (
	"TakeOut/router/middleware"
	"gorm.io/gorm"
	"reflect"
)

type Employee struct {
	ID       int    `json:"id,string"` // 通过json标签，使数据输出给前端时以json后面的名字为字段名
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	IdNumber string `json:"idNumber"`
	Status   int    `json:"status"`
	Universal
}

type EmployeeList []Employee

func (e *Employee) IsEmpty() bool {
	return reflect.DeepEqual(e, &Employee{})
}

func (e *Employee) BeforeCreate(*gorm.DB) (err error) {
	e.CreateUser = middleware.CurrentId
	e.UpdateUser = middleware.CurrentId
	return
}

func (e *Employee) BeforeUpdate(*gorm.DB) (err error) {
	e.UpdateUser = middleware.CurrentId
	return
}

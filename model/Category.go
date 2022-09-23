package model

import (
	"TakeOut/router/middleware"
	"gorm.io/gorm"
	"reflect"
)

type Category struct {
	ID   int    `json:"id,string"`
	Type int    `json:"type"`
	Name string `json:"name"`
	Sort int    `json:"sort"`
	Universal
}

type CategoryList []Category

func (c *Category) IsEmpty() bool {
	return reflect.DeepEqual(c, &Category{})
}

func (c *Category) BeforeCreate(*gorm.DB) (err error) {
	c.CreateUser = middleware.CurrentId
	c.UpdateUser = middleware.CurrentId
	return
}

func (c *Category) BeforeUpdate(*gorm.DB) (err error) {
	c.UpdateUser = middleware.CurrentId
	return
}

func GetCategoryByType(type_ string) CategoryList {
	var list CategoryList
	if type_ != "" {
		db.Order("sort asc, update_time desc").Where("type = ?", type_).Find(&list)
	} else {
		db.Order("sort asc, update_time desc").Find(&list)
	}
	return list
}

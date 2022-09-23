package v1

import (
	"TakeOut/api/v1/R"
	"TakeOut/model"
	"TakeOut/router/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddressBooks(c *gin.Context) {
	var addressBooks model.AddressBookList
	model.List(&addressBooks, "is_default desc, update_time desc", []string{"user_id"}, middleware.CurrentId)
	R.Success(c, addressBooks)
}

func AddAddressBook(c *gin.Context) {
	var addressBook model.AddressBook
	_ = c.ShouldBindJSON(&addressBook)
	model.Insert(addressBook)
	R.Success(c, "地址添加成功")
}

func GetAddressBookById(c *gin.Context) {
	id := c.Param("id")
	var data model.AddressBook
	model.GetOne(&data, []string{"id"}, id)
	R.Success(c, data)
}

func UpdateAddressBook(c *gin.Context) {
	var data model.AddressBook
	_ = c.ShouldBindJSON(&data)
	model.Update(data)
	R.Success(c, "11111")
}

func DeleteAddressBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("ids"))
	model.DeleteByIds(model.AddressBook{}, id)
	R.Success(c, "")
}

func GetDefault(c *gin.Context) {
	var data model.AddressBook
	model.GetOne(&data, []string{"user_id", "is_default"}, middleware.CurrentId, 1)
	if data.IsEmpty() {
		R.Error(c, "没有找到该对象！")
	} else {
		R.Success(c, data)
	}
}

func UpdateDefault(c *gin.Context) {
	var data model.AddressBook
	_ = c.ShouldBindJSON(&data)
	model.UpdateWith(model.AddressBook{}, "is_default", 1, map[string]any{"is_default": 0})
	data.IsDefault = 1
	model.Update(data)
	R.Success(c, "1111")
}

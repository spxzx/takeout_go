package v1

import (
	"TakeOut/api/v1/R"
	"TakeOut/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

// CategoryPage 分页菜品/套餐分类
func CategoryPage(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	var categories model.CategoryList
	var total int64
	model.Page(page, pageSize, nil, &categories, &total, "sort asc, update_time desc")
	R.Success(c, map[string]any{
		"records": categories,
		"total":   total,
	})
}

// AddCategory 添加菜品/套餐分类
func AddCategory(c *gin.Context) {
	var data, category model.Category
	_ = c.ShouldBindJSON(&data)
	model.GetOne(&category, []string{"name"}, data.Name)
	if !category.IsEmpty() {
		R.Error(c, "分类添加失败，分类"+category.Name+"已存在！")
		return
	}
	model.Insert(data)
	R.Success(c, "分类添加成功！")
}

// UpdateCategory 更新菜品/套餐分类信息
func UpdateCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	model.Update(data)
	R.Success(c, "分类更新成功")
}

// DeleteCategory 删除菜品/套餐分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	model.DeleteByIds(model.Category{}, id)
	R.Success(c, "分类删除成功")
}

func CateGoryList(c *gin.Context) {
	type_ := c.Query("type")
	R.Success(c, model.GetCategoryByType(type_))
}

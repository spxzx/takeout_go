package v1

import (
	"TakeOut/api/v1/R"
	"TakeOut/model"
	"TakeOut/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func DishDtoPage(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	name := c.Query("name")
	searchMap := make(map[string]string)
	if name != "" {
		searchMap["name"] = name
	}
	var dishes model.DishList
	var total int64
	model.Page(page, pageSize, searchMap, &dishes, &total, "update_time desc") // 获取目前dish分页内容
	var dishDtoList model.DishDtoList
	for _, dish := range dishes {
		var category model.Category
		model.GetOne(&category, []string{"id"}, dish.CategoryId)
		dishDto := model.DishDto{
			Dish:         dish,
			CategoryName: category.Name,
		}
		dishDtoList = append(dishDtoList, dishDto)
	}
	R.Success(c, map[string]any{
		"records": dishDtoList,
		"total":   total,
	})
}

func AddDish(c *gin.Context) {
	var dishDto model.DishDto
	_ = c.ShouldBindJSON(&dishDto)
	var dish model.Dish
	model.GetOne(&dish, []string{"name"}, dishDto.Name)
	if !dish.IsEmpty() {
		R.Error(c, "菜品添加失败，菜品"+dishDto.Name+"已存在！")
		return
	}
	model.InsertWithList(dishDto.Dish, dishDto.Flavors)
	R.Success(c, "添加菜品成功！")
}

func GetDishById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	R.Success(c, model.GetDishWithFlavors(id))
}

func UpdateDish(c *gin.Context) {
	var dto model.DishDto
	_ = c.ShouldBindJSON(&dto)
	model.UpdateDishWithFlavors(dto)
	R.Success(c, "菜品修改成功！")
}

func UpdateDishStatusToStop(c *gin.Context) {
	model.UpdateByIds(model.Dish{}, utils.StringsToInts(c.Query("ids")), map[string]any{"status": 0})
	R.Success(c, "菜品停售成功！")
}

func UpdateDishStatusToStart(c *gin.Context) {
	model.UpdateByIds(model.Dish{}, utils.StringsToInts(c.Query("ids")), map[string]any{"status": 1})
	R.Success(c, "菜品起售成功！")
}

func DeleteDish(c *gin.Context) {
	ids := utils.StringsToInts(c.Query("ids"))
	model.DeleteByIds(model.Dish{}, ids...)
	model.DeleteBatch(model.DishFlavor{}, "dish_id", ids)
	R.Success(c, "菜品删除成功！")
}

func DishList(c *gin.Context) {
	categoryId := c.Query("categoryId")
	status := c.Query("status")
	var dishes model.DishList
	model.List(&dishes, "sort asc, update_time desc", []string{"category_id", "status"}, categoryId, status)
	var dishDtos model.DishDtoList
	for _, dish := range dishes {
		var flavors model.DishFlavorList
		model.List(&flavors, "", []string{"dish_id"}, dish.ID)
		dishDtos = append(dishDtos, model.DishDto{
			Dish:    dish,
			Flavors: flavors,
		})
	}
	R.Success(c, dishDtos)
}

package v1

import (
	"TakeOut/api/v1/R"
	"TakeOut/model"
	"TakeOut/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func SetmealPage(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	name := c.Query("name")
	searchMap := make(map[string]string)
	if name != "" {
		searchMap["name"] = name
	}
	var setmeals model.SetmealList
	var total int64
	model.Page(page, pageSize, searchMap, &setmeals, &total, "update_time desc")
	var setmealDtoList model.SetmealDtoList
	for _, setmeal := range setmeals {
		var category model.Category
		model.GetOne(&category, []string{"id"}, setmeal.CategoryId)
		setmealDto := model.SetmealDto{
			Setmeal:      setmeal,
			CategoryName: category.Name,
		}
		setmealDtoList = append(setmealDtoList, setmealDto)
	}
	R.Success(c, map[string]any{
		"records": setmealDtoList,
		"total":   total,
	})
}

func AddSetmeal(c *gin.Context) {
	var setmealDto model.SetmealDto
	_ = c.ShouldBindJSON(&setmealDto)
	var setmeal model.Setmeal
	model.GetOne(&setmeal, []string{"name"}, setmealDto.Name)
	if !setmeal.IsEmpty() {
		R.Error(c, "套餐添加失败，菜品"+setmealDto.Name+"已存在！")
		return
	}
	model.InsertWithList(setmealDto.Setmeal, setmealDto.SetmealDishes)
	R.Success(c, "套餐添加成功！")
}

func GetSetmealById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	R.Success(c, model.GetSetmealWithDishes(id))
}

func UpdateSetmeal(c *gin.Context) {
	var dto model.SetmealDto
	_ = c.ShouldBindJSON(&dto)
	model.UpdateSetmealWithDishes(dto)
	R.Success(c, "套餐修改成功！")
}

func UpdateSetmealStatusToStop(c *gin.Context) {
	model.UpdateByIds(model.Setmeal{}, utils.StringsToInts(c.Query("ids")), map[string]any{"status": 0})
	R.Success(c, "套餐停售成功！")
}

func UpdateSetmealStatusToStart(c *gin.Context) {
	model.UpdateByIds(model.Setmeal{}, utils.StringsToInts(c.Query("ids")), map[string]any{"status": 1})
	R.Success(c, "套餐起售成功！")
}

func DeleteSetmeal(c *gin.Context) {
	ids := utils.StringsToInts(c.Query("ids"))
	model.DeleteByIds(model.Setmeal{}, ids...)
	model.DeleteBatch(model.SetmealDish{}, "setmeal_id", ids)
	R.Success(c, "套餐删除成功！")
}

func SetmealList(c *gin.Context) {
	cid := c.Query("categoryId")
	status := c.Query("status")
	var setmealList model.SetmealList
	model.List(&setmealList, "update_time desc", []string{"category_id", "status"}, cid, status)
	R.Success(c, setmealList)
}

func GetDishDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var setmealDishes model.SetmealDishList
	model.List(&setmealDishes, "", []string{"setmeal_id"}, id)
	var setmealDishDtos model.SetmealDishDtoList
	for _, setmealDish := range setmealDishes {
		setmealDishDto := model.SetmealDishDto{
			SetmealDish: setmealDish,
		}
		var dish model.Dish
		model.GetOne(&dish, []string{"id"}, setmealDish.DishId)
		setmealDishDto.Image = dish.Image
		setmealDishDto.Description = dish.Description
		setmealDishDtos = append(setmealDishDtos, setmealDishDto)
	}
	R.Success(c, setmealDishDtos)
}

package v1

import (
	"TakeOut/api/v1/R"
	"TakeOut/model"
	"TakeOut/router/middleware"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func UserOrderPage(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	var orders model.OrdersList
	var total int64
	model.PageWhere(page, pageSize, &orders, &total, "order_time desc",
		[]string{"user_id"}, middleware.CurrentId)
	var list model.OrderDtoList
	for _, order := range orders {
		var orderDetails model.OrderDetailList
		model.List(&orderDetails, "number asc", []string{"order_id"}, order.ID)
		sumNum := 0
		if len(orderDetails) > 0 {
			for _, od := range orderDetails {
				sumNum += od.Number
			}
		}
		dto := model.OrderDto{
			Orders:       order,
			OrderDetails: orderDetails,
			SunNum:       sumNum,
		}
		list = append(list, dto)
	}
	R.Success(c, map[string]any{
		"records": list,
		"total":   total,
	})
}

func SubmitOrder(c *gin.Context) {
	var order model.Orders
	_ = c.ShouldBindJSON(&order)
	userId := middleware.CurrentId
	var shoppingCartList model.ShoppingCartList
	model.List(&shoppingCartList, "", []string{"user_id"}, userId)
	if len(shoppingCartList) == 0 {
		R.Error(c, "购物车为空，不能下单！")
	}
	var user model.User
	model.GetOne(&user, []string{"id"}, userId)
	addressId := order.AddressBookId
	var addressBook model.AddressBook
	model.GetOne(&addressBook, []string{"id"}, addressId)
	if addressBook.IsEmpty() {
		R.Error(c, "用户地址信息有误，无法下单！")
	}
	var orderDetails model.OrderDetailList
	sum := 0.0
	for _, item := range shoppingCartList {
		orderDetail := model.OrderDetail{
			Number:     item.Number,
			DishFlavor: item.DishFlavor,
			DishId:     item.DishId,
			SetmealId:  item.SetmealId,
			Name:       item.Name,
			Image:      item.Image,
			Amount:     item.Amount,
		}
		sum += item.Amount
		orderDetails = append(orderDetails, orderDetail)
	}
	order.Status = 2
	order.Amount = sum
	order.UserId = userId
	order.Number = strconv.Itoa(int(time.Now().UnixNano()))
	order.UserName = user.Name
	order.Consignee = addressBook.Consignee
	order.Phone = addressBook.Phone
	order.Address = addressBook.Detail
	model.Insert(order)
	model.InsertBatch(orderDetails)
	model.DeleteBatch(model.ShoppingCart{}, "user_id", middleware.CurrentId)
	R.Success(c, "success")
}

func OrderAgain(c *gin.Context) {
	var order model.Orders
	_ = c.ShouldBindJSON(&order)
	var orderDetails model.OrderDetailList
	model.List(&orderDetails, "", []string{"order_id"}, order.ID)
	for _, item := range orderDetails {
		shoppingCart := model.ShoppingCart{
			Name:       item.Name,
			Number:     item.Number,
			Amount:     item.Amount,
			DishFlavor: item.DishFlavor,
		}
		if item.DishId != 0 {
			var dish model.Dish
			model.GetOne(&dish, []string{"id"}, item.DishId)
			shoppingCart.Image = dish.Image
			shoppingCart.DishId = item.DishId
		} else {
			var setmeal model.Setmeal
			model.GetOne(&setmeal, []string{"id"}, item.SetmealId)
			shoppingCart.Image = setmeal.Image
			shoppingCart.SetmealId = item.SetmealId
		}
		model.Insert(shoppingCart)
	}
	R.Success(c, "再来一单！")
}

func OrderPage(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	number := c.Query("number")
	searchMap := make(map[string]string)
	if number != "" {
		searchMap["number"] = number
	}
	var orders model.OrdersList
	var total int64
	model.Page(page, pageSize, searchMap, &orders, &total, "order_time desc")
	var ordersDtos model.OrderDtoList
	for _, item := range orders {
		dto := model.OrderDto{
			Orders:   item,
			Username: item.UserName,
			Address:  item.Address,
		}
		ordersDtos = append(ordersDtos, dto)
	}
	R.Success(c, map[string]any{
		"records": ordersDtos,
		"total":   total,
	})
}

func UpdateOrderStatus(c *gin.Context) {
	var order model.Orders
	_ = c.ShouldBindJSON(&order)
	model.UpdateWith(order, "id", order.ID, map[string]any{"status": order.Status})
	R.Success(c, "成功！")
}

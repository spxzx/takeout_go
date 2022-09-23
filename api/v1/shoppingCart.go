package v1

import (
	"TakeOut/api/v1/R"
	"TakeOut/model"
	"TakeOut/router/middleware"
	"github.com/gin-gonic/gin"
)

func AddToShoppingCart(c *gin.Context) {
	var shoppingCart, cart model.ShoppingCart
	_ = c.ShouldBindJSON(&shoppingCart)
	shoppingCart.UserId = middleware.CurrentId
	dishId := shoppingCart.DishId
	if dishId != 0 {
		model.GetOne(&cart, []string{"user_id", "dish_id"}, middleware.CurrentId, dishId)
	} else {
		model.GetOne(&cart, []string{"user_id", "setmeal_id"}, middleware.CurrentId, shoppingCart.SetmealId)
	}
	if !cart.IsEmpty() {
		number := cart.Number
		cart.Number = number + 1
		model.Update(cart)
	} else {
		shoppingCart.Number = 1
		model.Insert(shoppingCart)
		cart = shoppingCart
	}
	R.Success(c, cart)
}

func SubFromToShoppingCart(c *gin.Context) {
	var shoppingCart, cart model.ShoppingCart
	_ = c.ShouldBindJSON(&shoppingCart)
	model.GetOne(&cart, []string{"dish_id"}, shoppingCart.DishId)
	number := cart.Number
	if number > 1 {
		cart.Number = number - 1
		model.Update(cart)
	} else {
		model.DeleteByIds(cart, cart.ID)
	}
	R.Success(c, cart)
}

func ShoppingCartList(c *gin.Context) {
	var list model.ShoppingCartList
	model.List(&list, "create_time desc", []string{"user_id"}, middleware.CurrentId)
	R.Success(c, list)
}

func CleanShoppingCartList(c *gin.Context) {
	model.DeleteBatch(model.ShoppingCart{}, "user_id", middleware.CurrentId)
	R.Success(c, "success")
}

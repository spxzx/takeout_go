package router

import (
	v1 "TakeOut/api/v1"
	"TakeOut/router/middleware"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	store := cookie.NewStore([]byte("ah2i5j"))
	r.Use(sessions.Sessions("session", store))

	r.StaticFS("/backend", http.Dir("./web/backend"))
	r.StaticFS("/front", http.Dir("./web/front"))

	r.GET("/backend", func(c *gin.Context) {
		c.Redirect(302, "/backend/page/login/login.html")
	})
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/front/page/login.html")
	})
	r.POST("/employee/login", v1.Login)
	r.POST("/user/sendMsg", v1.SendMessage)
	r.POST("/user/login", v1.UserLogin)

	// 上面的请求不用经过验证
	r.Use(middleware.Authorize())

	common := r.Group("/common")
	{
		common.POST("/upload", v1.Upload)
		common.GET("/download", v1.Download)
	}

	employee := r.Group("/employee")
	{
		employee.POST("/logout", v1.Logout)
		employee.GET("/page", v1.EmployeePage)
		employee.POST("", v1.AddEmployee)
		employee.GET("/:id", v1.GetEmployeeById)
		employee.PUT("", v1.UpdateEmployee)
	}

	category := r.Group("/category")
	{
		category.GET("/page", v1.CategoryPage)
		category.POST("", v1.AddCategory)
		category.PUT("", v1.UpdateCategory)
		category.DELETE("", v1.DeleteCategory)
		category.GET("/list", v1.CateGoryList)
	}

	dish := r.Group("/dish")
	{
		dish.GET("/page", v1.DishDtoPage)
		dish.POST("", v1.AddDish)
		dish.GET("/:id", v1.GetDishById)
		dish.PUT("", v1.UpdateDish)
		dish.POST("/status/0", v1.UpdateDishStatusToStop)
		dish.POST("/status/1", v1.UpdateDishStatusToStart)
		dish.DELETE("", v1.DeleteDish)
		dish.GET("/list", v1.DishList)
	}

	setmeal := r.Group("/setmeal")
	{
		setmeal.GET("/page", v1.SetmealPage)
		setmeal.POST("", v1.AddSetmeal)
		setmeal.GET("/:id", v1.GetSetmealById)
		setmeal.PUT("", v1.UpdateSetmeal)
		setmeal.POST("/status/0", v1.UpdateSetmealStatusToStop)
		setmeal.POST("/status/1", v1.UpdateSetmealStatusToStart)
		setmeal.DELETE("", v1.DeleteSetmeal)
		setmeal.GET("/list", v1.SetmealList)
		setmeal.GET("/dish/:id", v1.GetDishDetail)
	}

	order := r.Group("/order")
	{
		order.GET("/userPage", v1.UserOrderPage)
		order.POST("/submit", v1.SubmitOrder)
		order.POST("/again", v1.OrderAgain)
		order.GET("/page", v1.OrderPage)
		order.PUT("", v1.UpdateOrderStatus)
	}

	user := r.Group("/user")
	user.POST("/loginout", v1.UserLogout)

	addressBook := r.Group("/addressBook")
	{
		addressBook.GET("/list", v1.AddressBooks)
		addressBook.POST("", v1.AddAddressBook)
		addressBook.GET("/:id", v1.GetAddressBookById)
		addressBook.PUT("", v1.UpdateAddressBook)
		addressBook.DELETE("", v1.DeleteAddressBook)
		addressBook.GET("/default", v1.GetDefault)
		addressBook.PUT("/default", v1.UpdateDefault)
	}

	shoppingCart := r.Group("/shoppingCart")
	{
		shoppingCart.POST("/add", v1.AddToShoppingCart)
		shoppingCart.POST("/sub", v1.SubFromToShoppingCart)
		shoppingCart.GET("/list", v1.ShoppingCartList)
		shoppingCart.DELETE("/clean", v1.CleanShoppingCartList)
	}

	if err := r.Run(""); err != nil {
		log.Fatal(err)
	}

}

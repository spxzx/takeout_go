package middleware

import (
	"TakeOut/api/v1/R"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

var CurrentId int

// Authorize 验证用户是否登录中间件
func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		eid, uid := session.Get("employee"), session.Get("user")
		if eid == nil && uid == nil {
			log.Println("用户未登录...")
			R.Error(c, "NOTLOGIN")
			c.Abort()
		}
		if eid != nil {
			log.Printf("用户已登录, ID为 %v, 可放行...\n", eid)
			CurrentId = eid.(int)
		} else if uid != nil {
			log.Printf("用户已登录, ID为 %v, 可放行...\n", uid)
			CurrentId = uid.(int)
		}
		c.Next()
	}
}

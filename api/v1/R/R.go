package R

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
	})
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": data,
	})
}

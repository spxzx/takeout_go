package v1

import (
	"TakeOut/api/v1/R"
	"github.com/gin-gonic/gin"
)

const dst = "./web/images/"

func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	if err := c.SaveUploadedFile(file, dst+file.Filename); err != nil {
		R.Error(c, "图片上传失败，请稍后重试！")
	}
	R.Success(c, file.Filename)
}

func Download(c *gin.Context) {
	c.File(dst + c.Query("name"))
}

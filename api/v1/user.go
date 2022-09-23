package v1

import (
	"TakeOut/api/v1/R"
	"TakeOut/model"
	"TakeOut/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func SendMessage(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	phone := user.Phone
	if phone == "" {
		R.Error(c, "短信发送失败，请稍后重新尝试！")
		return
	}
	code := utils.GenerateValidateCode()
	session := sessions.Default(c)
	session.Set("phone", code)
	if err := session.Save(); err != nil {
		R.Error(c, "短信发送失败，请稍后重新尝试！")
		return
	}
	log.Println("[[[[[ 验证码 " + code + " ]]]]]")
	R.Success(c, "手机验证码发送成功！")
	return
}

func UserLogin(c *gin.Context) {
	var m map[string]string
	_ = c.ShouldBindJSON(&m)
	phone, code := m["phone"], m["code"]
	session := sessions.Default(c)
	codeInSession := session.Get("phone")
	if codeInSession != "" && codeInSession == code {
		var user model.User
		model.GetOne(&user, []string{"phone"}, phone)
		if user.IsEmpty() {
			user.Phone = phone
			user.Status = 1
			model.Insert(user)
		}
		session.Set("user", user.ID)
		if err := session.Save(); err != nil {
			R.Error(c, "登陆失败，请稍后重试！")
			return
		}
		R.Success(c, user)
		return
	}
	R.Error(c, "登录失败，请稍后重新尝试！")
}

func UserLogout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("user")
	if err := session.Save(); err != nil {
		R.Error(c, "退出失败！")
		return
	}
	R.Success(c, "退出成功！")
}

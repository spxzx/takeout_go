package v1

import (
	"TakeOut/api/v1/R"
	"TakeOut/model"
	"TakeOut/utils/md5"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Login 登录
func Login(c *gin.Context) {
	var data, employee model.Employee
	_ = c.ShouldBindJSON(&data)
	model.GetOne(&employee, []string{"username"}, data.Username)
	if employee.IsEmpty() || employee.Password != md5.Transfer(data.Password) {
		R.Error(c, "用户名或密码错误，请检查后重新输入！")
		return
	}
	if employee.Status == 0 {
		R.Error(c, "账号已被封禁！")
		return
	}
	session := sessions.Default(c)
	session.Set("employee", employee.ID)
	if err := session.Save(); err != nil {
		R.Error(c, "登陆失败，请稍后重试！")
		return
	}
	R.Success(c, data)
}

// Logout 登出
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("employee")
	if err := session.Save(); err != nil {
		R.Error(c, "退出失败！")
		return
	}
	R.Success(c, "退出成功！")
}

// EmployeePage 员工页面分页展示
func EmployeePage(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	name := c.Query("name")
	searchMap := make(map[string]string)
	if name != "" {
		searchMap["name"] = name // 数据库    { 字段名 : 字段值 }
	}
	var employees model.EmployeeList
	var total int64
	model.Page(page, pageSize, searchMap, &employees, &total, "update_time desc")
	R.Success(c, map[string]any{
		"records": employees,
		"total":   total,
	})
}

// AddEmployee 增加员工
func AddEmployee(c *gin.Context) {
	var data, employee model.Employee
	_ = c.ShouldBindJSON(&data)
	model.GetOne(&employee, []string{"username"}, data.Username)
	if !employee.IsEmpty() {
		R.Error(c, "员工添加失败，员工"+employee.Username+"已存在！")
		return
	}
	data.Status = 1
	data.Password = md5.Transfer("123456")
	model.Insert(data)
	R.Success(c, "员工添加成功！")
}

// GetEmployeeById 根据id获取员工
func GetEmployeeById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var employee model.Employee
	model.GetOne(&employee, []string{"id"}, id)
	R.Success(c, employee)
}

// UpdateEmployee 更新员工信息/状态
func UpdateEmployee(c *gin.Context) {
	var data model.Employee
	_ = c.ShouldBindJSON(&data)
	updateMap := make(map[string]any)
	if data.Name == "" {
		updateMap["id"] = data.ID
		updateMap["status"] = data.Status
	}
	model.Update(data, updateMap)
	R.Success(c, "编辑员工信息/状态成功！")
}

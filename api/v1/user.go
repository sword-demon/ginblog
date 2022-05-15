package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
 * @describe:用户模块的控制器
 */

var code int

// UserExist 查询用户是否存在
func UserExist(c *gin.Context) {

}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var userModel model.User
	_ = c.ShouldBindJSON(&userModel)
	code = model.CheckUser(userModel.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&userModel)
	}
	if code == errmsg.ErrorUsernameUsed {
		code = errmsg.ErrorUsernameUsed
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    userModel,
	})
}

// 查询单个用户

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	users, total := model.GetUsers(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"data":    gin.H{"users": users, "total": total},
	})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {

}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {

}

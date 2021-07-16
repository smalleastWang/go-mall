package controllers

import (
	"github.com/gin-gonic/gin"
	"go-mall/middleware"
	"go-mall/models"
	"go-mall/utils"
)

func Login(c *gin.Context) {
	var query models.LoginBo
	if err := c.ShouldBindJSON(&query); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}

	if hasSession := middleware.HasSession(c); hasSession == true {
		utils.OkWithMessage("用户已登陆", c)
		return
	}

	user := models.UserDetailByName(query.Username)

	if err := utils.Compare(user.Password, query.Password); err != nil {
		utils.FailWithMessage("用户名或密码错误", c)
		return
	}

	middleware.SaveAuthSession(c, user.ID)
	utils.OkWithMessage("登录成功", c)
}

func Logout(c *gin.Context) {
	if hasSession := middleware.HasSession(c); hasSession == false {
		utils.FailWithMessage("用户未登陆", c)
		return
	}
	middleware.ClearAuthSession(c)
	utils.OkWithMessage("退出成功", c)
}

func Register(c *gin.Context) {
	//models.UserDelete(6)
	var query models.RegisterBo
	if err := c.ShouldBindJSON(&query); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var user models.User
	_ = utils.CopyProperties(&user, query)

	if hasSession := middleware.HasSession(c); hasSession == true {
		utils.FailWithMessage("用户已登陆", c)
		return
	}

	if existUser := models.UserDetailByName(user.Username); existUser.ID != 0 {
		utils.FailWithMessage("用户名已存在", c)
		return
	}


	if query.Password != query.ConfirmPassword {
		utils.FailWithMessage("密码不一致", c)
		return
	}

	if pwd, err := utils.Encrypt(query.Password); err == nil {
		query.Password = pwd
	}
	models.AddUser(&user)
	user = models.UserDetailByName(query.Username)

	middleware.SaveAuthSession(c, user.ID)

	utils.OkWithMessage("注册成功", c)

}

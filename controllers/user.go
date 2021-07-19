package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mall/middleware"
	"go-mall/models"
	"go-mall/service"
	"go-mall/utils"
)

// @Tags user
// @Summary 登录
// @accept application/json
// @Produce application/json
// @Param data body models.LoginBo true "客户用户名, 客户手机号码"
// @Success 200 {object} models.LoginVo
// @Router /public/user/login [post]
func Login(c *gin.Context) {
	var query models.LoginBo
	if err := c.ShouldBindJSON(&query); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}

	user := service.UserDetailByName(query.Username)

	if err := utils.Compare(user.Password, query.Password); err != nil {
		utils.FailWithMessage("用户名或密码错误", c)
		return
	}

	token, err := middleware.GenToken(user.Username)
	if err != nil {
		utils.FailWithMessage("Token生成失败", c)
		return
	}
	// 组装返回数据
	var loginVo models.LoginVo
	_ = utils.CopyProperties(&loginVo, user)
	loginVo.Token = token
	utils.OkWithData(loginVo, c)
}

// @Tags user
// @Summary 退出登录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"退出成功"}"
// @Router /public/user/logout [post]
func Logout(c *gin.Context) {
	//if hasSession := middleware.HasSession(c); hasSession == false {
	//	utils.FailWithMessage("用户未登陆", c)
	//	return
	//}
	//middleware.ClearAuthSession(c)
	utils.OkWithMessage("退出成功", c)
}
// @Tags user
// @Summary 注册
// @accept application/json
// @Produce application/json
// @Param data body models.RegisterBo true "哈哈"
// @Success 200 {object}  models.LoginVo
// @Router /public/user/register [post]
func Register(c *gin.Context) {
	var query models.RegisterBo
	if err := c.ShouldBindJSON(&query); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	var user models.User
	_ = utils.CopyProperties(&user, query)


	if existUser := service.UserDetailByName(user.Username); existUser.ID != 0 {
		utils.FailWithMessage("用户名已存在", c)
		return
	}

	if query.Password != query.ConfirmPassword {
		utils.FailWithMessage("密码不一致", c)
		return
	}

	if pwd, err := utils.Encrypt(query.Password); err == nil {
		user.Password = pwd
	}
	service.AddUser(&user)
	//user = models.UserDetailByName(query.Username)
	if user.ID == 0 {
		utils.FailWithMessage("注册失败", c)
		return
	}
	token, err := middleware.GenToken(user.Username)
	if err != nil {
		utils.FailWithMessage("Token生成失败", c)
		return
	}
	var loginVo models.LoginVo
	_ = utils.CopyProperties(&loginVo, user)
	loginVo.Token = token

	utils.OkWithData(loginVo, c)

}
// @Tags user
// @Summary 修改密码
// @accept application/json
// @Produce application/json
// @Param data body models.ResetPasswordBo true "修改密码"
// @Success 200 {string} string
// @Router /public/user/reset-password [post]
func ResetPassword(c *gin.Context) {
	//models.UserDelete(6)
	var query models.ResetPasswordBo
	if err := c.ShouldBindJSON(&query); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	if pwd, err := utils.Encrypt(query.NewPassword); err == nil {
		query.NewPassword = pwd
	}
	existUser := service.UserDetail(query.Id)
	if existUser.ID == 0 {
		utils.FailWithMessage("该用户不存在", c)
		return
	}
	if err := utils.Compare(existUser.Password, query.OldPassword); err != nil {
		utils.FailWithMessage("原密码错误", c)
		return
	}

	result := service.UserUpdate(&query)

	fmt.Println(result)
	utils.OkWithMessage("密码修改成功", c)

}

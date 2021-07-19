package service

import (
	"go-mall/models"
	"go-mall/utils"
)

// 用户注册时新增
func AddUser(user *models.User) {
	utils.GetDB().Create(&user)
	return
}

func UserDetailByName(name string) (user models.User) {
	utils.GetDB().Where("username = ?", name).First(&user)
	return
}

func UserDetailByEmail(email string) (user models.User) {
	utils.GetDB().Where("email = ?", email).First(&user)
	return
}

func UserDetail(id uint) (user models.User) {
	utils.GetDB().Where("id = ?", id).First(&user)
	return
}
func UserUpdate(data *models.ResetPasswordBo) (user models.User) {
	utils.GetDB().Where("id = ?", data.Id).First(&user).Update("password", data.NewPassword)
	return
}

func UserDelete(id uint) (user models.User) {
	utils.GetDB().Delete(&models.User{}, id)
	return
}

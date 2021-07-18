package models

import (
	"go-mall/utils"
)

type User struct {
	OrmDefaultModel
	Username	string 	`gorm:"username"`
	Password	string 	`gorm:"password"`
	Email		string 	`gorm:"email"`
	Phone		int 	`gorm:"phone"`
	Nickname	string 	`gorm:"nickname"`
	Avatar		string 	`gorm:"avatar"`
}

type LoginBo struct {
	Username	string `json:"username" form:"username" binding:"required"`
	Password	string `json:"password" form:"password" binding:"required"`
}

type RegisterBo struct {
	Username		string `json:"username" form:"username" binding:"required"`
	Password		string `json:"password" form:"password" binding:"required"`
	ConfirmPassword	string `json:"confirmPassword" form:"confirmPassword" binding:"required"`
	Phone			int `json:"phone" form:"phone" binding:"required"`
	Email			string `json:"email" form:"email" binding:"required"`
	Nickname	string `json:"nickname" form:"nickname"`
}
type ResetPasswordBo struct {
	Id 	uint `json:"id"`
	OldPassword string `json:"oldPassword" form:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" form:"newPassword" binding:"required"`
}


// 用户注册时新增
func AddUser(user *User) {
	utils.GetDB().Create(&user)
	return
}

func UserDetailByName(name string) (user User) {
	utils.GetDB().Where("username = ?", name).First(&user)
	return
}

func UserDetailByEmail(email string) (user User) {
	utils.GetDB().Where("email = ?", email).First(&user)
	return
}

func UserDetail(id uint) (user User) {
	utils.GetDB().Where("id = ?", id).First(&user)
	return
}
func UserUpdate(data *ResetPasswordBo) (user User) {
	utils.GetDB().Where("id = ?", data.Id).First(&user).Update("password", data.NewPassword)
	return
}

func UserDelete(id uint) (user User) {
	utils.GetDB().Delete(&User{}, id)
	return
}
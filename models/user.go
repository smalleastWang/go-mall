package models

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

type LoginVo struct {
	ID			uint `json:"userId"`
	Token 		string `json:"token"`
	Username 	string `json:"username"`
	Nickname	string `json:"nickname"`
	Avatar		string `json:"avatar"`
}
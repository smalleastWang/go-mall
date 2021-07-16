package models


type Address struct {
	OrmDefaultModel
	AddressBo
	//Name		string 	`gorm:"name"`
	//UserId 		uint 	`gorm:"user_id"`
	//Mobile		string 	`gorm:"mobile"`
	//Province	string 	`gorm:"province"`
	//City		string 	`gorm:"city"`
	//County		string 	`gorm:"county"`
	//Address		string 	`gorm:"address"`
}

type AddressBo struct {
	Name		string 	`gorm:"name" json:"name" binding:"required"`
	UserId 		uint 	`gorm:"user_id" json:"user_id" binding:"required"`
	Mobile		string 	`gorm:"mobile" json:"mobile" binding:"required"`
	Province	string 	`gorm:"province" json:"province" binding:"required"`
	City		string 	`gorm:"city" json:"city" binding:"required"`
	County		string 	`gorm:"county" json:"county" binding:"required"`
	Address		string 	`gorm:"address" json:"address" binding:"required"`
}
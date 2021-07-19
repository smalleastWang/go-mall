package models

type Goods struct {
	OrmDefaultModel
	Cid			string 	`json:"cid" gorm:"cid" binding:"required"`
	Price			float64 `json:"price" gorm:"price" binding:"required"`
	DiscountPrice	float64 `json:"discountPrice" gorm:"discount_price" binding:"required"`
	Name			string 	`json:"name" gorm:"name" binding:"required"`
	Desc			string 	`json:"desc" gorm:"desc" binding:"required"`
	Remark			string 	`json:"remark" gorm:"remark" binding:"required"`
}

type GoodsParams struct {
	PageBo
	Goods
}
package service

import (
	"go-mall/models"
	"go-mall/utils"
)

func AddGoods(goods *models.Goods) {
	utils.GetDB().Create(&goods)
	return
}
func GetGoodsList(goodsParams *models.GoodsParams) (result models.PageVo) {
	result.Data =  make([]string, 0)
	offset := (goodsParams.Page - 1) * goodsParams.PageSize
	result.Page = goodsParams.Page
	result.PageSize = goodsParams.PageSize

	goodsDb := utils.GetDB().Where(&goodsParams)
	var total int64
	goodsDb.Count(&total)
	var goodsList []models.Goods
	goodsDb.Offset(offset).Limit(goodsParams.PageSize).Find(&goodsList)
	result.Total = int(total)
	if goodsList != nil {
		result.Data = goodsList
	}
	return
}
func GetGoodsById(id uint) (goods models.Goods) {
	utils.GetDB().Where("id = ?", id).First(&goods)
	return
}

func UpdateGoodsById(data *models.Goods) (goods models.Goods) {
	utils.GetDB().Where("id = ?", data.ID).First(&goods).Updates(data)
	return
}

func DeleteGoods(id uint) (user models.Goods) {
	utils.GetDB().Delete(&models.Goods{}, id)
	return
}

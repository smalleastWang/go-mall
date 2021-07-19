package controllers

import (
	"github.com/gin-gonic/gin"
	"go-mall/models"
	"go-mall/service"
	"go-mall/utils"
	"strconv"
)


// @Tags 商品相关
// @Summary 分页查询商品列表
// @accept application/json
// @Produce application/json
// @Param data body models.GoodsParams true "哈哈"
// @Success 200 {array}  models.Goods
// @Router /api/goods [get]
func GetGoodsList (c *gin.Context)  {
	var query models.GoodsParams
	if err := c.ShouldBindJSON(&query); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	result := service.GetGoodsList(&query)
	utils.OkWithData(result, c)
}
// @Tags 商品相关
// @Summary 分页查询商品列表
// @accept application/json
// @Produce application/json
// @Param id param {string} true "哈哈"
// @Success 200 {object}  models.Goods
// @Router /api/goods/:id [post]
func GetGoodsById (c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	result := service.GetGoodsById(uint(id))
	if result.ID == 0 {
		utils.FailWithMessage("没有查询到数据", c)
	} else {
		utils.OkWithData(result, c)
	}

}
// @Tags 商品相关
// @Summary 分页查询商品列表
// @accept application/json
// @Produce application/json
// @Param id param {string} true "哈哈"
// @Success 200 {object}  models.Goods
// @Router /api/goods/:id [post]
func CreateGoods (c *gin.Context)  {
	var goods models.Goods
	if err := c.ShouldBindJSON(&goods); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	service.AddGoods(&goods)
	utils.OkWithMessage("添加成功", c)
}

func UpdateGoods (c *gin.Context)  {

}

func DeleteGoods (c *gin.Context)  {

}

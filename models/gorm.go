package models

import "gorm.io/gorm"

type OrmDefaultModel struct {
	gorm.Model
	ID			uint `json:"id" gorm:"id"`
}

type PageVo struct {
	Data			interface{} `json:"data"`
	Page			int `json:"page"`
	PageSize		int `json:"pageSize"`
	Total			int `json:"total"`
}
type PageBo struct {
	Page			int
	PageSize		int
}

package models

import "gorm.io/gorm"

type OrmDefaultModel struct {
	gorm.Model
	ID			uint `gorm:"id"`
}

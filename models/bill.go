package models

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	ID          uint
	Name        string     `gorm:"type:varchar(255); NOT NULL" json:"name" binding:"required"`
	Amount      uint32     `gorm:"NOT NULL" json:"amount" binding:"required"`
	Category    []Category `gorm:"NA" json:"category"`
	Description string     `gorm:"default:description for good" json:"description"`
}

type Category struct {
	gorm.Model
	ID     uint16
	Name   string
	BillID uint
}

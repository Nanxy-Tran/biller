package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID     uint16
	Name   string `gorm:"NOT NULL" binding:"required" json:"name"`
	BillID uint
}

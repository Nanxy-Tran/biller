package models

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	ID          uint
	Name        string
	Amount      uint32
	Category    []Category
	Description string `gorm:"default:description for good"`
}

type Category struct {
	gorm.Model
	ID     uint16
	Name   string
	BillID uint
}

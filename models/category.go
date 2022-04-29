package models

type Category struct {
	ID     uint16
	Name   string `gorm:"NOT NULL" binding:"required" json:"name"`
	BillID uint
}

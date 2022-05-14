package models

type Tag struct {
	Base
	Name       string  `gorm:"not null; unique" binding:"required" json:"name"`
	UserID     uint    `gorm:"default:null"`
	CategoryID uint    `gorm:"not null"`
	Bills      []*Bill `gorm:"many2many:bills_tags;"`
}

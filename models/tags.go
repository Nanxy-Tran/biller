package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name   string  `gorm:"not null; unique" binding:"required" json:"name"`
	UserID uint    `gorm:"default:null"`
	Bills  []*Bill `gorm:"many2many:bills_tags;"`
}

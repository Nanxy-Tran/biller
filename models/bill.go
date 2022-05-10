package models

import "gorm.io/gorm"

type Bill struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null"`
	Amount uint32 `json:"amount"`
	Tags   []*Tag `json:"tags" gorm:"many2many:bill_tags;"`
}

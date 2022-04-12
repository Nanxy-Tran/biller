package models

import "time"

type Bill struct {
	ID          uint
	Name        string    `json:"name" binding:"required"`
	Amount      uint32    `json:"amount" binding:"required"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

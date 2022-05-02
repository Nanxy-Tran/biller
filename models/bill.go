package models

import "time"

type Bill struct {
	ID          uint
	Name        string    `json:"name"`
	Amount      uint32    `json:"amount"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

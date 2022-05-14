package models

type Category struct {
	Base
	Name  string `json:"name" gorm:"index:priority, unique"`
	Color string `json:"color"`
}

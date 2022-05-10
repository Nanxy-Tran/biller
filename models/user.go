package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"username,omitempty" form:"username"`
	Password []byte `json:"password,omitempty" form:"password"`
	Email    string `json:"email,omitempty" form:"email" gorm:"unique"`
	Role     string `json:"role,omitempty"`
	Tags     []Tag  `json:"tags,omitempty"`
}

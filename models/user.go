package models

type User struct {
	Base
	UserName string `json:"username,omitempty" form:"username"`
	Password []byte `json:"password,omitempty" form:"password"`
	Email    string `json:"email,omitempty" form:"email" gorm:"unique"`
	Role     string `json:"role,omitempty"`
	Tags     []Tag  `json:"tags,omitempty"`
}

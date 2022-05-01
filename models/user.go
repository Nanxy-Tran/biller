package models

type User struct {
	ID       uint   `json:"ID,omitempty"`
	UserName string `json:"username,omitempty" form:"username"`
	Password []byte `json:"password,omitempty" form:"password"`
	Email    string `json:"email,omitempty" form:"email"`
	Role     string `json:"role,omitempty"`
}

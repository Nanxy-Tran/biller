package models

type User struct {
	ID       uint   `json:"ID,omitempty"`
	UserName string `json:"user_name,omitempty" form:"user_name"`
	Password []byte `json:"password,omitempty" form:"password"`
	Email    string `json:"email,omitempty" form:"email"`
	Role     string `json:"role,omitempty"`
}

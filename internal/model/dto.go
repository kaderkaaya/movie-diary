package model

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=32" gorm:"unique"`
	Email    string `json:"email"    binding:"required,email" gorm:"unique"`
	Password string `json:"password" binding:"required,min=6"`
}

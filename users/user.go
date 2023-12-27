package user

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	RegisterNumber int32  `gorm:"column:register_number" json:"register_number"`
	Email          string `gorm:"column:email" json:"email"`
	Password       string `gorm:"column:password" json:"password"`
	Nickname       string `gorm:"column:nickname" json:"nickname"`
	Company        string `gorm:"column:company" json:"company"`
	IsActive       string `gorm:"column:is_active" json:"is_active"`
	LoggedIn       string `gorm:"column:loggedin" json:"loggedin"`
}

func (Users) TableName() string {
	return "users"
}

type UserRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponseDto struct {
	RegisterNumber int32  `json:"register_number"`
	Email          string `json:"email"`
	Nickname       string `json:"nickname"`
	Company        string `json:"company"`
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
}

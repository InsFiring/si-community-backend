package user

import (
	"time"

	"gorm.io/gorm"
)

type Tokens struct {
	gorm.Model
	TokenId        int64     `gorm:"column:token_id" json:"token_id"`
	RegisterNumber int32     `gorm:"column:register_number" json:"register_number"`
	Token          string    `gorm:"column:token" json:"token"`
	ExpirationTime time.Time `gorm:"column:expiration_time" json:"expiration_time"`
}

func (Tokens) TableName() string {
	return "user_tokens"
}

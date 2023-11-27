package models

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

// type ArticleReplies struct {
// 	gorm.Model
// 	ReplyId    int32     `gorm:"column:reply_id" json:"reply_id"`
// 	ArticleId  int32     `gorm:"column:article_id" json:"article_id"`
// 	Nickname   string    `gorm:"column:nickname" json:"nickname"`
// 	Contents   string    `gorm:"column:contents" json:"contents"`
// 	Likes      int32     `gorm:"column:likes" json:"likes"`
// 	Unlikes    int32     `gorm:"column:unlikes" json:"unlikes"`
// 	isModified string    `gorm:"column:is_modified" json:"is_modified"`
// 	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
// 	ModifiedAt time.Time `gorm:"column:modified_at" json:"modified_at"`
// }

// func (ArticleReplies) TableName() string {
// 	return "article_replies"
// }

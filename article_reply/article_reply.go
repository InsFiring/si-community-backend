package article_reply

import "gorm.io/gorm"

type ArticleReplies struct {
	gorm.Model
	ReplyId       int32  `gorm:"primaryKey;column:reply_id" json:"reply_id"`
	ArticleId     int32  `gorm:"column:article_id" json:"article_id"`
	ParentReplyId int32  `gorm:"column:parent_reply_id" json:"parent_reply_id"`
	Nickname      string `gorm:"column:nickname" json:"nickname"`
	Contents      string `gorm:"column:contents" json:"contents"`
	Likes         int32  `gorm:"column:likes" json:"likes"`
	Unlikes       int32  `gorm:"column:unlikes" json:"unlikes"`
	IsModified    string `gorm:"column:is_modified" json:"is_modified"`
}

func (ArticleReplies) TableName() string {
	return "article_replies"
}

type ArticleReplyRequestDto struct {
	ArticleId     int32  `gorm:"column:article_id" json:"article_id"`
	ParentReplyId int32  `gorm:"column:parent_reply_id" json:"parent_reply_id,omitempty"`
	Nickname      string `gorm:"column:nickname" json:"nickname"`
	Contents      string `gorm:"column:contents" json:"contents"`
}

type ArticleReplyModifyDto struct {
	Contents string `gorm:"column:contents" json:"contents"`
}

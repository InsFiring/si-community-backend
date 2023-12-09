package article

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	ArticleId  int32  `gorm:"column:article_id" json:"article_id"`
	Nickname   string `gorm:"column:nickname" json:"nickname"`
	Company    string `gorm:"column:company" json:"company"`
	Ratings    int32  `gorm:"column:ratings" json:"ratings"`
	Title      string `gorm:"column:title" json:"title"`
	Contents   string `gorm:"column:contents" json:"contents"`
	ViewCounts int32  `gorm:"column:view_counts" json:"view_counts"`
	Likes      int32  `gorm:"column:likes" json:"likes"`
	Unlikes    int32  `gorm:"column:unlikes" json:"unlikes"`
	IsModified string `gorm:"column:is_modified" json:"is_modified"`
}

func (Article) TableName() string {
	return "articles"
}

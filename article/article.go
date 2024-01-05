package article

import "gorm.io/gorm"

type Articles struct {
	gorm.Model
	ArticleId  int32  `gorm:"primaryKey;column:article_id" json:"article_id"`
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

func (Articles) TableName() string {
	return "articles"
}

type ArticleRequestDto struct {
	Ratings  int32  `json:"ratings"`
	Title    string `json:"title"`
	Contents string `json:"contents"`
	Nickname string `json:"nickname"`
	Company  string `json:"company"`
}

type ArticleModifyDto struct {
	ArticleId int32  `json:"article_id"`
	Ratings   int32  `json:"ratings"`
	Title     string `json:"title"`
	Contents  string `json:"contents"`
}

type ArticleSearchDto struct {
	Ratings  *int32 `json:"ratings,omitempty"`
	Title    string `json:"title,omitempty"`
	Contents string `json:"contents,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Company  string `json:"company,omitempty"`
}

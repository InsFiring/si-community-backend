package article

import (
	"fmt"

	"gorm.io/gorm"
)

const True string = "y"
const False string = "n"
const DefaultCount = 0

type ArticleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{db}
}

func (r *ArticleRepository) AddArticle(articleRequestDto ArticleRequestDto) (Articles, error) {
	fmt.Println("ArticleRepository AddArticle")

	article := Articles{
		Nickname:   articleRequestDto.Nickname,
		Company:    articleRequestDto.Company,
		Ratings:    articleRequestDto.Ratings,
		Title:      articleRequestDto.Title,
		Contents:   articleRequestDto.Contents,
		ViewCounts: DefaultCount,
		Likes:      DefaultCount,
		Unlikes:    DefaultCount,
		IsModified: False,
	}

	return article, r.db.Omit("article_id").Create(&article).Error
}

func (r *ArticleRepository) GetArticles() []Articles {
	fmt.Println("ArticleRepository GetArticles")

	var articles []Articles
	r.db.Table("articles").Find(&articles)

	return articles
}

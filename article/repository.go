package article

import (
	"errors"
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

func (r *ArticleRepository) GetArticleById(articleId int32) (Articles, error) {
	fmt.Println("ArticleRepository GetArticleById")

	var article Articles
	var count int64

	result := r.db.Table("articles").
		Where(&Articles{ArticleId: articleId}).
		Find(&article)

	result.Count(&count)

	err := result.Error
	if err != nil {
		return article, err
	}

	if count == 0 {
		return article, errors.New("글이 없습니다.")
	}

	article.ViewCounts += 1

	result.Update("view_counts", article.ViewCounts)

	return article, nil
}

func (r *ArticleRepository) ModifyArticle(articleModifyDto ArticleModifyDto) (Articles, error) {
	fmt.Println("ArticleRepository ModifyArticle")

	var article Articles
	var count int64

	result := r.db.Table("articles").
		Where(&Articles{ArticleId: articleModifyDto.ArticleId}).
		Find(&article)

	result.Count(&count)

	err := result.Error
	if err != nil {
		return article, err
	}

	if count == 0 {
		return article, errors.New("글이 없습니다.")
	}

	article.Ratings = articleModifyDto.Ratings
	article.Title = articleModifyDto.Title
	article.Contents = articleModifyDto.Contents
	article.IsModified = True

	r.db.Updates(article)

	return article, nil
}

func (r *ArticleRepository) PlusLike(articleId int32) (Articles, error) {
	fmt.Println("ArticleRepository PlusLike")

	var article Articles
	var count int64

	result := r.db.Table("articles").
		Where(&Articles{ArticleId: articleId}).
		Find(&article)

	result.Count(&count)

	err := result.Error
	if err != nil {
		return article, err
	}

	if count == 0 {
		return article, errors.New("글이 없습니다.")
	}

	article.Likes += 1

	result.Update("likes", article.Likes)

	return article, nil
}

func (r *ArticleRepository) CancelLike(articleId int32) (Articles, error) {
	fmt.Println("ArticleRepository PlusLike")

	var article Articles
	var count int64

	result := r.db.Table("articles").
		Where(&Articles{ArticleId: articleId}).
		Find(&article)

	result.Count(&count)

	err := result.Error
	if err != nil {
		return article, err
	}

	if count == 0 {
		return article, errors.New("글이 없습니다.")
	}

	article.Likes -= 1

	result.Update("likes", article.Likes)

	return article, nil
}

func (r *ArticleRepository) PlusUnlike(articleId int32) (Articles, error) {
	fmt.Println("ArticleRepository PlusUnlike")

	var article Articles
	var count int64

	result := r.db.Table("articles").
		Where(&Articles{ArticleId: articleId}).
		Find(&article)

	result.Count(&count)

	err := result.Error
	if err != nil {
		return article, err
	}

	if count == 0 {
		return article, errors.New("글이 없습니다.")
	}

	article.Unlikes += 1

	result.Update("unlikes", article.Unlikes)

	return article, nil
}

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

func (r *ArticleRepository) GetArticles(page int, offset int) []Articles {
	fmt.Println("ArticleRepository GetArticles")

	var articles []Articles

	r.db.Table("articles").Offset((page - 1) * offset).
		Limit(offset).
		Find(&articles)

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

func (r *ArticleRepository) ModifyArticle(articleModifyDto ArticleModifyDto, articleId int32) (Articles, error) {
	fmt.Println("ArticleRepository ModifyArticle")

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

func (r *ArticleRepository) CancelUnlike(articleId int32) (Articles, error) {
	fmt.Println("ArticleRepository CancelUnlike")

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

	article.Unlikes -= 1

	result.Update("unlikes", article.Unlikes)

	return article, nil
}

func (r *ArticleRepository) SearchArticles(articleSearchDto ArticleSearchDto, page int, offset int) ([]Articles, error) {
	fmt.Println("ArticleRepository SearchArticles")

	var articles []Articles
	var result *gorm.DB

	if articleSearchDto.Ratings != nil {
		result = r.db.Table("articles").
			Where(&Articles{Ratings: *articleSearchDto.Ratings}).
			Offset((page - 1) * offset).
			Limit(offset).
			Find(&articles)
	} else if articleSearchDto.Title != "" {
		result = r.db.Table("articles").
			Where("title Like ?", "%"+articleSearchDto.Title+"%").
			Offset((page - 1) * offset).
			Limit(offset).
			Find(&articles)
	} else if articleSearchDto.Contents != "" {
		result = r.db.Table("articles").
			Where("contents Like ?", "%"+articleSearchDto.Contents+"%").
			Offset((page - 1) * offset).
			Limit(offset).
			Find(&articles)
	} else if articleSearchDto.Nickname != "" {
		result = r.db.Table("articles").
			Where("nickname Like ?", "%"+articleSearchDto.Nickname+"%").
			Offset((page - 1) * offset).
			Limit(offset).
			Find(&articles)
	} else if articleSearchDto.Company != "" {
		result = r.db.Table("articles").
			Where("company Like ?", "%"+articleSearchDto.Company+"%").
			Offset((page - 1) * offset).
			Limit(offset).
			Find(&articles)
	}

	if err := result.Error; result.Error != nil {
		return articles, err
	}

	return articles, nil
}

func (r *ArticleRepository) DeleteArticle(articleId int32) (int32, error) {
	fmt.Println("ArticleRepository DeleteArticle")

	var article Articles
	var count int64

	result := r.db.Table("articles").
		Where(&Articles{ArticleId: articleId}).
		Find(&article)

	result.Count(&count)

	err := result.Error
	if err != nil {
		return articleId, err
	}

	if count == 0 {
		return articleId, errors.New("글이 없습니다.")
	}

	return articleId, r.db.Where(&Articles{ArticleId: articleId}).Delete(&Articles{}).Error
}

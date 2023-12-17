package articlereply

import (
	"errors"

	"gorm.io/gorm"
)

const True string = "y"
const False string = "n"
const DefaultCount = 0

type ArticleReplyRepository struct {
	db *gorm.DB
}

func NewArticleReplyRepository(db *gorm.DB) *ArticleReplyRepository {
	return &ArticleReplyRepository{db}
}

func (r *ArticleReplyRepository) AddArticleReply(articleReplyRequestDto ArticleReplyRequestDto) (ArticleReplies, error) {
	articleReply := ArticleReplies{
		ArticleId:  articleReplyRequestDto.ArticleId,
		Nickname:   articleReplyRequestDto.Nickname,
		Contents:   articleReplyRequestDto.Contents,
		Likes:      DefaultCount,
		Unlikes:    DefaultCount,
		IsModified: False,
	}

	return articleReply, r.db.Omit("reply_id").Create(&articleReply).Error
}

func (r *ArticleReplyRepository) GetArticleRepliesByArticleId(articleId int32) ([]ArticleReplies, error) {
	var articleReplies []ArticleReplies
	var count int64

	result := r.db.Table("article_replies").
		Where(&ArticleReplies{ArticleId: articleId}).
		Find(&articleReplies)

	result.Count(&count)

	err := result.Error
	if err != nil {
		return articleReplies, err
	}

	if count == 0 {
		return articleReplies, errors.New("댓글이 없습니다.")
	}

	return articleReplies, nil
}

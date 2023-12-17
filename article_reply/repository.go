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

func (r *ArticleReplyRepository) ModifyArticleReply(articleReplyModifyDto ArticleReplyModifyDto) (ArticleReplies, error) {
	var articleReply ArticleReplies
	var count int64

	result := r.db.Table("article_replies").
		Where(&ArticleReplies{
			ArticleId: articleReplyModifyDto.ArticleId,
			ReplyId:   articleReply.ReplyId}).
		Find(&articleReply)

	result.Count(&count)

	err := result.Error
	if err != nil {
		return articleReply, err
	}

	if count == 0 {
		return articleReply, errors.New("댓글이 없습니다.")
	}

	articleReply.Contents = articleReplyModifyDto.Contents
	articleReply.IsModified = True

	r.db.Updates(articleReply)

	return articleReply, nil
}

func (r *ArticleReplyRepository) PlusReplyLike(articleId int32, replyId int32) (ArticleReplies, error) {
	var articleReplies ArticleReplies
	var count int64

	result := r.db.Table("article_replies").
		Where(&ArticleReplies{
			ArticleId: articleId,
			ReplyId:   replyId}).
		Find(&articleReplies)

	result.Count(&count)

	err := result.Error
	if err != nil {
		return articleReplies, err
	}

	if count == 0 {
		return articleReplies, errors.New("댓글이 없습니다.")
	}

	articleReplies.Likes += 1

	result.Update("likes", articleReplies.Likes)

	return articleReplies, nil
}

func (r *ArticleReplyRepository) CancelReplyLike(articleId int32, replyId int32) (ArticleReplies, error) {
	var articleReplies ArticleReplies
	var count int64

	result := r.db.Table("article_replies").
		Where(&ArticleReplies{
			ArticleId: articleId,
			ReplyId:   replyId}).
		Find(&articleReplies)

	result.Count(&count)

	err := result.Error
	if err != nil {
		return articleReplies, err
	}

	if count == 0 {
		return articleReplies, errors.New("댓글이 없습니다.")
	}

	articleReplies.Likes -= 1

	result.Update("likes", articleReplies.Likes)

	return articleReplies, nil
}

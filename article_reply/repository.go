package articlereply

import "gorm.io/gorm"

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

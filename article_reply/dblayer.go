package articlereply

type Dblayer interface {
	AddArticleReply(articleReplyRequestDto ArticleReplyRequestDto) (ArticleReplies, error)
	GetArticleRepliesByArticleId(articleId int32) (ArticleReplies, error)
}

package article

type Dblayer interface {
	AddArticle(articleRequestDto ArticleRequestDto) (Articles, error)
	GetArticles() ([]Articles, error)
	GetArticleById(articleId int32) (Articles, error)
	ModifyArticle(articleModifyDto ArticleModifyDto)
	PlusLike(articleId int32) (Articles, error)
	CancelLike(articleId int32) (Articles, error)
	PlusUnlike(articleId int32) (Articles, error)
	CancelUnlike(articleId int32) (Articles, error)
	DeleteArticle(articleId int32) (int32, error)
}

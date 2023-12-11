package article

type Dblayer interface {
	AddArticle(articleRequestDto ArticleRequestDto) (Articles, error)
	GetArticles() ([]Articles, error)
	GetArticleById(articleId int32) (Articles, error)
	ModifyArticle(articleModifyDto ArticleModifyDto)
}

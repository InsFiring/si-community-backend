package article

type Dblayer interface {
	AddArticle(articleRequestDto ArticleRequestDto) (Articles, error)
	GetArticles() ([]Articles, error)
}

package article

type Dblayer interface {
	AddArticle(articleRequestDto ArticleRequestDto) (Articles, error)
}

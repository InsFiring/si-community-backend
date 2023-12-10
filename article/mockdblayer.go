package article

import "encoding/json"

type MockDBLayer struct {
	err      error
	articles []Articles
}

func NewMockDBLayer(
	articles []Articles) *MockDBLayer {
	return &MockDBLayer{
		articles: articles,
	}
}

func NewMockDBLayerWithData() *MockDBLayer {
	ARTICLES := `[
		{
			"article_id" : 1,
			"nickname" : "test",
			"company" : "keke",
			"ratings" : 5,
			"title" : "이건제목",
			"contents" : "글 내용입니다.",
			"view_counts" : 0,
			"likes" : 0,
			"unlikes" : 0,
			"is_modified" : "n",
			"created_at" : "2023-12-11 00:07:52",
			"updated_at" : "2023-12-11 00:07:52",
			"deleted_at" : null
		},
		{
			"article_id" : 2,
			"nickname" : "test",
			"company" : "keke",
			"ratings" : 5,
			"title" : "이건제목2",
			"contents" : "글 내용입니다.222",
			"view_counts" : 0,
			"likes" : 0,
			"unlikes" : 0,
			"is_modified" : "n",
			"created_at" : "2023-12-11 00:08:52",
			"updated_at" : "2023-12-11 00:08:52",
			"deleted_at" : null
		}
	]`

	var articles []Articles
	json.Unmarshal([]byte(ARTICLES), &articles)

	return NewMockDBLayer(articles)
}

func (mock *MockDBLayer) GetMockArticleData() []Articles {
	return mock.articles
}

func (mock *MockDBLayer) SetError(err error) {
	mock.err = err
}

func (mock *MockDBLayer) AddUser(articleRequestDto ArticleRequestDto) ([]Articles, error) {

	if mock.err != nil {
		return []Articles{}, mock.err
	}

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

	mock.articles = append(mock.articles, article)
	return mock.articles, nil
}

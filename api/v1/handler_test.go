package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"si-community/article"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddArticle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockDBLayer := article.NewMockDBLayerWithData()
	handler, err := NewHandler()
	if err != nil {
		fmt.Println(err)
	}

	r := gin.New()
	r.POST("/v1/articles", handler.AddArticle)

	testArticle := article.ArticleRequestDto{
		Ratings:  3,
		Title:    "이건제목3",
		Contents: "글 내용입니다.3",
		Nickname: "test",
		Company:  "keke",
	}

	jsonData, err := json.Marshal(testArticle)
	assert.NoError(t, err)

	request, err := http.NewRequest(http.MethodPost, "/v1/articles", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)

	request.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, request)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response article.Articles
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedResult := article.Articles{
		Nickname:   "test",
		Company:    "keke",
		Ratings:    3,
		Title:      "이건제목3",
		Contents:   "글 내용입니다.3",
		ViewCounts: 0,
		Likes:      0,
		Unlikes:    0,
		IsModified: "n",
	}

	assert.Equal(t, expectedResult, response)

	testArticles, err := mockDBLayer.AddUser(testArticle)
	assert.NoError(t, err)
	assert.Equal(t, len(testArticles), 3)
}

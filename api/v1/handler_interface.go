package api

import (
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	ChangePassword(c *gin.Context)
	AddArticle(c *gin.Context)
	GetArticles(c *gin.Context)
	GetArticleById(c *gin.Context)
	ModifyArticle(c *gin.Context)
}

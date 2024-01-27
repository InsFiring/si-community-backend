package api

import (
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	ChangePassword(c *gin.Context)
	HasEmail(c *gin.Context)
	ChangeUserInfo(c *gin.Context)
	AddArticle(c *gin.Context)
	GetArticles(c *gin.Context)
	GetArticleById(c *gin.Context)
	ModifyArticle(c *gin.Context)
	PlusLike(c *gin.Context)
	CancelLike(c *gin.Context)
	PlusUnlike(c *gin.Context)
	CancelUnlike(c *gin.Context)
	SearchArticles(c *gin.Context)
	DeleteArticle(c *gin.Context)
	AddArticleReply(c *gin.Context)
	GetArticleRepliesByArticleId(c *gin.Context)
	ModifyArticleReply(c *gin.Context)
	PlusReplyLike(c *gin.Context)
	CancelReplyLike(c *gin.Context)
	PlusReplyUnlike(c *gin.Context)
	CancelReplyUnlike(c *gin.Context)
	DeleteArticleReply(c *gin.Context)
}

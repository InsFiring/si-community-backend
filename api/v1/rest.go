package api

import (
	docs "si-community/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunAPI(address string) error {
	handler, err := NewHandler()
	if err != nil {
		return err
	}

	return RunApiWithHandler(address, handler)
}

func RunApiWithHandler(address string, handler HandlerInterface) error {
	r := gin.Default()

	const BASEPATH = "/v1"

	usersGroup := r.Group(BASEPATH + "/users")
	{
		// 유저 추가
		usersGroup.POST("", handler.AddUser)
		usersGroup.POST("/signin", handler.SignIn)
		usersGroup.POST("/changePassword", handler.ChangePassword)
	}

	r.POST(BASEPATH+"/article", handler.AddArticle)
	r.GET(BASEPATH+"/articles", handler.GetArticles)
	r.GET(BASEPATH+"/article/:id", handler.GetArticleById)
	r.PUT(BASEPATH+"/article/:id", handler.ModifyArticle)
	r.GET(BASEPATH+"/article/:id/like", handler.PlusLike)
	r.GET(BASEPATH+"/article/:id/cancel_like", handler.CancelLike)
	r.GET(BASEPATH+"/article/:id/unlike", handler.PlusUnlike)
	r.GET(BASEPATH+"/article/:id/cancel_unlike", handler.CancelUnlike)
	r.DELETE(BASEPATH+"/article/:id", handler.DeleteArticle)

	r.POST(BASEPATH+"/article_reply", handler.AddArticleReply)
	r.GET(BASEPATH+"/article/:id/article_replies", handler.GetArticleRepliesByArticleId)
	r.PUT(BASEPATH+"/article/:id/article_replies/:reply_id", handler.ModifyArticleReply)
	r.GET(BASEPATH+"/article/:id/article_replies/:reply_id/like", handler.PlusReplyLike)

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r.Run(address)
}

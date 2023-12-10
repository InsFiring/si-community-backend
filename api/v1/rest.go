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

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r.Run(address)
}

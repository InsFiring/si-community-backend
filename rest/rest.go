package rest

import "github.com/gin-gonic/gin"

func RunAPI(address string) error {
	handler, err := NewHandler()
	if err != nil {
		return err
	}

	return RunApiWithHandler(address, handler)
}

func RunApiWithHandler(address string, handler HandlerInterface) error {
	r := gin.Default()

	usersGroup := r.Group("/users")
	{
		// 유저 추가
		usersGroup.POST("", handler.AddUser)
		usersGroup.POST("/signin", handler.SignIn)
		usersGroup.POST("/changePassword", handler.ChangePassword)
	}

	return r.Run(address)
}

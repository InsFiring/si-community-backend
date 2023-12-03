package rest

import (
	"fmt"
	"net/http"
	"si-community/dblayer"
	"si-community/models"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	ChangePassword(c *gin.Context)
}

type Handler struct {
	db dblayer.DBlayer
}

func NewHandler() (*Handler, error) {
	dbConn, err := dblayer.DBConnection()
	if err != nil {
		fmt.Println("DBConn error")
		return nil, err
	}
	handler := new(Handler)
	handler.db = dbConn
	return handler, nil
}

func NewHandlerWithDB(db dblayer.DBlayer) HandlerInterface {
	return &Handler{db: db}
}

// @Description 자세한 설명은 이곳에 적습니다.
// @name add user
// @Accept  json
// @Produce  json
// @Router /v1/users [post]
// @Success 200 {object} models.UserResponseDto
func (h *Handler) AddUser(c *gin.Context) {
	if h.db == nil {
		fmt.Println("DB is nil")
		return
	}

	var user models.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("ShouldBindJSON error")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.db.AddUser(user)
	if err != nil {
		fmt.Println("db add user error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *Handler) SignIn(c *gin.Context) {
	if h.db == nil {
		fmt.Println("DB is nil")
		return
	}

	var userRequestDto models.UserRequestDto
	err := c.ShouldBindJSON(&userRequestDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponseDto, err := h.db.SignInUser(userRequestDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userResponseDto)

	return
}

// TODO: 이메일 인증 부분 구현 필요(일단은 인증 없이 구현)ㄴ
func (h *Handler) ChangePassword(c *gin.Context) {
	if h.db == nil {
		fmt.Println("DB is nil")
		return
	}

	var userRequestDto models.UserRequestDto
	err := c.ShouldBindJSON(&userRequestDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponseDto, err := h.db.ChangePassword(userRequestDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userResponseDto)

	return

}

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

// @Description 회원 가입
// @name Add User
// @Accept  json
// @Produce  json
// @Param users body models.Users false "email, password, nickname, company만 있으면 됨"
// @Router /v1/users [post]
// @Success 201 {object} models.Users
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

// @Description 로그인
// @name Sign In
// @Accept  json
// @Produce  json
// @Param users body models.UserRequestDto true "로그인 input"
// @Router /v1/users/signin [post]
// @Success 201 {object} models.UserResponseDto
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

// @Description 비밀번호 변경
// @name ChangePassword
// @Accept  json
// @Produce  json
// @Param users body models.UserRequestDto true "비밀번호 변경 input"
// @Router /v1/users/changePassword [post]
// @Success 201 {object} models.UserResponseDto
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

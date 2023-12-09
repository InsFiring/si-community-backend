package api

import (
	"fmt"
	"net/http"
	"si-community/config"
	user "si-community/users"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userRepsitory user.UserRepository
}

func addTestData(userRepository user.UserRepository) {
	user := user.Users{
		Email:    "test@gmail.com",
		Password: "test1234",
		Nickname: "test",
		Company:  "keke",
	}

	userRepository.AddUser(user)
}

func NewHandler() (*Handler, error) {
	dbConn, err := config.DBConnection()
	if err != nil {
		fmt.Println("DBConn error")
		return nil, err
	}
	handler := new(Handler)
	handler.userRepsitory = *user.NewUserRepository(dbConn)

	addTestData(handler.userRepsitory)
	return handler, nil
}

// @Description 회원 가입
// @name Add User
// @Accept  json
// @Produce  json
// @Param users body user.Users false "email, password, nickname, company만 있으면 됨"
// @Router /v1/users [post]
// @Success 201 {object} user.Users
func (h *Handler) AddUser(c *gin.Context) {
	var user user.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println("ShouldBindJSON error")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userRepsitory.AddUser(user)
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
// @Param users body user.Users true "로그인 input"
// @Router /v1/users/signin [post]
// @Success 201 {object} user.Users
func (h *Handler) SignIn(c *gin.Context) {
	var userRequestDto user.UserRequestDto
	err := c.ShouldBindJSON(&userRequestDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponseDto, err := h.userRepsitory.SignInUser(userRequestDto)
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
// @Param users body user.Users true "비밀번호 변경 input"
// @Router /v1/users/changePassword [post]
// @Success 201 {object} user.Users
func (h *Handler) ChangePassword(c *gin.Context) {
	var userRequestDto user.UserRequestDto
	err := c.ShouldBindJSON(&userRequestDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponseDto, err := h.userRepsitory.ChangePassword(userRequestDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, userResponseDto)

	return

}

func (h *Handler) AddArticle(c *gin.Context) {
	return
}

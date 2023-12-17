package api

import (
	"fmt"
	"net/http"
	"si-community/article"
	articlereply "si-community/article_reply"
	"si-community/config"
	user "si-community/users"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userRepsitory          user.UserRepository
	articleRepository      article.ArticleRepository
	articleReplyRepository articlereply.ArticleReplyRepository
}

func addTestUser(userRepository user.UserRepository) {
	user := user.Users{
		Email:    "test@gmail.com",
		Password: "test1234",
		Nickname: "test",
		Company:  "keke",
	}

	userRepository.AddUser(user)
}

func addTestArticle(articleRepository article.ArticleRepository) {
	articleRequestDto := article.ArticleRequestDto{
		Ratings:  5,
		Title:    "이건제목",
		Contents: "글 내용입니다.",
		Nickname: "test",
		Company:  "keke",
	}

	articles := articleRepository.GetArticles()
	if len(articles) == 0 {
		articleRepository.AddArticle(articleRequestDto)
	}
}

func addTestArticleReplies(articleReplyRepository articlereply.ArticleReplyRepository) {
	articleRepliyRequestDto1 := articlereply.ArticleReplyRequestDto{
		ArticleId: 1,
		Nickname:  "dori",
		Contents:  "괜찮은 글이군요.",
	}

	articleRepliyRequestDto2 := articlereply.ArticleReplyRequestDto{
		ArticleId: 1,
		Nickname:  "neo",
		Contents:  "도움이 많이 됐습니다 ㅎㅎ",
	}

	articleReplies, err := articleReplyRepository.GetArticleRepliesByArticleId(int32(1))
	if err != nil && len(articleReplies) == 0 {
		articleReplyRepository.AddArticleReply(articleRepliyRequestDto1)
		articleReplyRepository.AddArticleReply(articleRepliyRequestDto2)
	}
}

func NewHandler() (*Handler, error) {
	dbConn, err := config.DBConnection()
	if err != nil {
		fmt.Println("DBConn error")
		return nil, err
	}
	handler := new(Handler)
	handler.userRepsitory = *user.NewUserRepository(dbConn)
	handler.articleRepository = *article.NewArticleRepository(dbConn)
	handler.articleReplyRepository = *articlereply.NewArticleReplyRepository(dbConn)

	addTestUser(handler.userRepsitory)
	addTestArticle(handler.articleRepository)
	addTestArticleReplies(handler.articleReplyRepository)

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
	var articleRequestDto article.ArticleRequestDto
	err := c.ShouldBindJSON(&articleRequestDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := h.articleRepository.AddArticle(articleRequestDto)
	if err != nil {
		fmt.Println("db add article error")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, article)
}

// TODO: 페이지네이션 처리 필요
func (h *Handler) GetArticles(c *gin.Context) {
	articles := h.articleRepository.GetArticles()

	c.JSON(http.StatusOK, articles)
}

func (h *Handler) GetArticleById(c *gin.Context) {
	paramId := c.Param("id")
	articleId, err := strconv.Atoi(paramId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := h.articleRepository.GetArticleById(int32(articleId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (h *Handler) ModifyArticle(c *gin.Context) {
	var articleModifyDto article.ArticleModifyDto
	err := c.ShouldBindJSON(&articleModifyDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := h.articleRepository.ModifyArticle(articleModifyDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (h *Handler) PlusLike(c *gin.Context) {
	paramId := c.Param("id")
	articleId, err := strconv.Atoi(paramId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := h.articleRepository.PlusLike(int32(articleId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (h *Handler) CancelLike(c *gin.Context) {
	paramId := c.Param("id")
	articleId, err := strconv.Atoi(paramId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := h.articleRepository.CancelLike(int32(articleId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (h *Handler) PlusUnlike(c *gin.Context) {
	paramId := c.Param("id")
	articleId, err := strconv.Atoi(paramId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := h.articleRepository.PlusUnlike(int32(articleId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (h *Handler) CancelUnlike(c *gin.Context) {
	paramId := c.Param("id")
	articleId, err := strconv.Atoi(paramId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := h.articleRepository.CancelUnlike(int32(articleId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (h *Handler) DeleteArticle(c *gin.Context) {
	paramId := c.Param("id")
	articleId, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.articleRepository.DeleteArticle(int32(articleId))
	if err != nil {
		fmt.Println("Delete result error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) AddArticleReply(c *gin.Context) {
	var articleReplyRequestDto articlereply.ArticleReplyRequestDto

	err := c.ShouldBindJSON(&articleReplyRequestDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.articleReplyRepository.AddArticleReply(articleReplyRequestDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, articleReplyRequestDto)
}

func (h *Handler) GetArticleRepliesByArticleId(c *gin.Context) {
	paramId := c.Param("id")
	articleId, err := strconv.Atoi(paramId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articleReplies, err := h.articleReplyRepository.GetArticleRepliesByArticleId(int32(articleId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, articleReplies)
}

func (h *Handler) ModifyArticleReply(c *gin.Context) {
	var articleReplyModifyDto articlereply.ArticleReplyModifyDto
	err := c.ShouldBindJSON(&articleReplyModifyDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articleReply, err := h.articleReplyRepository.ModifyArticleReply(articleReplyModifyDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, articleReply)
}

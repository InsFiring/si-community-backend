package api

import (
	"fmt"
	"net/http"
	"si-community/article"
	articlereply "si-community/article_reply"
	user "si-community/users"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	userRepsitory          user.UserRepository
	articleRepository      article.ArticleRepository
	articleReplyRepository articlereply.ArticleReplyRepository
}

type EmailResponse struct {
	HasEmail bool `json:"has_email"`
}

func addTestUser(userRepository user.UserRepository) {
	user := user.Users{
		Email:    "test@gmail.com",
		Password: "test1234",
		Nickname: "test",
		Company:  "keke",
		IsAdmin:  "n",
	}

	userRepository.AddUser(user)
}

func addTestArticle(articleRepository article.ArticleRepository) {

	articles := articleRepository.GetArticles(1, 1)
	if len(articles) == 0 {
		for i := 1; i < 100; i++ {
			articleRequestDto := article.ArticleRequestDto{
				Ratings:  int32(i % 5),
				Title:    fmt.Sprintf("이건 제목 %d", i),
				Contents: fmt.Sprintf("이건 내용 %d", i),
				Nickname: "test",
				Company:  "keke",
			}
			articleRepository.AddArticle(articleRequestDto)
		}
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

func NewHandler(dbConn *gorm.DB) (*Handler, error) {
	handler := new(Handler)
	handler.userRepsitory = *user.NewUserRepository(dbConn)
	handler.articleRepository = *article.NewArticleRepository(dbConn)
	handler.articleReplyRepository = *articlereply.NewArticleReplyRepository(dbConn)

	addTestUser(handler.userRepsitory)
	addTestArticle(handler.articleRepository)
	addTestArticleReplies(handler.articleReplyRepository)

	return handler, nil
}

// @Tags users
// @Description 회원 가입
// @name Add User
// @Accept  json
// @Produce  json
// @Param Users body user.Users false "email, password, nickname, company만 있으면 됨 / is_admin 컬럼으로 회사 관리자인지 일반유저인지 구분 가능"
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

// @Tags users
// @Description 로그인
// @name Sign In
// @Accept  json
// @Produce  json
// @Param UserRequestDto body user.UserRequestDto true "로그인 input / new_password는 없어도 됨."
// @Router /v1/users/signin [post]
// @Success 200 {object} user.UserResponseDto
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
}

// @Tags users
// @Description 비밀번호 변경
// @name ChangePassword
// @Accept  json
// @Produce  json
// @Param UserRequestDto body user.UserRequestDto true "비밀번호 변경 input"
// @Router /v1/users/changePassword [post]
// @Success 200 {object} user.UserResponseDto
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

}

// @Tags users
// @Description 이메일 중복 체크
// @name HasEmail
// @Accept  json
// @Produce  json
// @Param UserRequestDto body user.UserRequestDto true "이메일값만 있으면 됨"
// @Router /v1/users/emails [post]
// @Success 200 {object} EmailResponse
func (h *Handler) HasEmail(c *gin.Context) {
	var userRequestDto user.UserRequestDto
	err := c.ShouldBindJSON(&userRequestDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hasEmail, err := h.userRepsitory.HasEmail(userRequestDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emailResponse := &EmailResponse{
		HasEmail: hasEmail,
	}

	c.JSON(http.StatusOK, emailResponse)
}

// @Tags accounts
// @Description 게시글 추가
// @name AddArticle
// @Accept  json
// @Produce  json
// @Param ArticleRequestDto body article.ArticleRequestDto true "ratings, title, contents, nickname, company 필수"
// @Router /v1/article [post]
// @Success 200 {object} article.Articles
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

// @Tags accounts
// @Description 게시글 조회
// @name GetArticles
// @Accept  json
// @Produce  json
// @Param page query string true "page 번호"
// @Param offset query string true "offset 숫자"
// @Router /v1/articles [get]
// @Success 200 {object} article.Articles
func (h *Handler) GetArticles(c *gin.Context) {
	paramPage := c.Query("page")
	paramOffset := c.Query("offset")

	page, err := strconv.Atoi(paramPage)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	offset, err := strconv.Atoi(paramOffset)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articles := h.articleRepository.GetArticles(page, offset)

	c.JSON(http.StatusOK, articles)
}

// @Tags accounts
// @Description 게시글 ID로 조회
// @name GetArticleById
// @Accept  json
// @Produce  json
// @Param id path int true "게시글 ID"
// @Router /v1/article/{id} [get]
// @Success 200 {object} article.Articles
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

// @Tags accounts
// @Description 게시글 수정
// @name ModifyArticle
// @Accept  json
// @Produce  json
// @Param ArticleModifyDto body article.ArticleModifyDto true "수정 관련 DTO 사용"
// @Router /v1/article/{id} [put]
// @Success 200 {object} article.Articles
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

// @Tags accounts
// @Description 게시글 좋아요 수 추가
// @name PlusLike
// @Accept  json
// @Produce  json
// @Param id path int true "게시글 ID"
// @Router /v1/article/{id}/like [get]
// @Success 200 {object} article.Articles
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

// @Tags accounts
// @Description 게시글 좋아요 취소
// @name CancelLike
// @Accept  json
// @Produce  json
// @Param id path int true "게시글 ID"
// @Router /v1/article/{id}/cancel_like [get]
// @Success 200 {object} article.Articles
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

// @Tags accounts
// @Description 게시글 싫어요 추가
// @name PlusUnlike
// @Accept  json
// @Produce  json
// @Param id path int true "게시글 ID"
// @Router /v1/article/{id}/unlike [get]
// @Success 200 {object} article.Articles
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

// @Tags accounts
// @Description 게시글 싫어요 취소
// @name PlusUnlike
// @Accept  json
// @Produce  json
// @Param id path int true "게시글 ID"
// @Router /v1/article/{id}/cancel_unlike [get]
// @Success 200 {object} article.Articles
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

// @Tags accounts
// @Description 게시글 삭제
// @name DeleteArticle
// @Accept  json
// @Produce  json
// @Param id path int true "게시글 ID"
// @Router /v1/article/{id} [delete]
// @Success 200 {object} article.Articles
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

func (h *Handler) PlusReplyLike(c *gin.Context) {
	paramId := c.Param("id")
	paramReplyId := c.Param("reply_id")

	articleId, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	replyId, err := strconv.Atoi(paramReplyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articleReply, err := h.articleReplyRepository.PlusReplyLike(int32(articleId), int32(replyId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, articleReply)
}

func (h *Handler) CancelReplyLike(c *gin.Context) {
	paramId := c.Param("id")
	paramReplyId := c.Param("reply_id")

	articleId, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	replyId, err := strconv.Atoi(paramReplyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articleReply, err := h.articleReplyRepository.CancelReplyLike(int32(articleId), int32(replyId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, articleReply)
}

func (h *Handler) PlusReplyUnlike(c *gin.Context) {
	paramId := c.Param("id")
	paramReplyId := c.Param("reply_id")

	articleId, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	replyId, err := strconv.Atoi(paramReplyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articleReply, err := h.articleReplyRepository.PlusReplyUnlike(int32(articleId), int32(replyId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, articleReply)
}

func (h *Handler) CancelReplyUnlike(c *gin.Context) {
	paramId := c.Param("id")
	paramReplyId := c.Param("reply_id")

	articleId, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	replyId, err := strconv.Atoi(paramReplyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	articleReply, err := h.articleReplyRepository.CancelReplyUnlike(int32(articleId), int32(replyId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, articleReply)
}

func (h *Handler) DeleteArticleReply(c *gin.Context) {
	paramId := c.Param("id")
	paramReplyId := c.Param("reply_id")

	articleId, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	replyId, err := strconv.Atoi(paramReplyId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deletedReplyId, err := h.articleReplyRepository.DeleteArticleReply(int32(articleId), int32(replyId))
	if err != nil {
		fmt.Println("Delete result error")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deletedReplyId)
}

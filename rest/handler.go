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

package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h handler) AddBook(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Added a Book successfully",})
}
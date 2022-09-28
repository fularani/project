package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h handler) DeleteBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Deleted a Book successfully",})
}
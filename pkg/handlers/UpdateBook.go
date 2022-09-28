package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h handler) UpdateBook(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Updated a Book successfully",})
}
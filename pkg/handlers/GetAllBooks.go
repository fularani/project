package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h handler) GetAllBooks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GOT all Books!",})
}
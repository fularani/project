package handlers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"project/models"
)

func (h handler) CreateQuery(c *gin.Context) {

	BOOK := new(models.Book)
	h.DB.Migrator().CreateTable(&BOOK)
    c.JSON(http.StatusOK, gin.H{"message": "Created a Book Table successfully",})
}
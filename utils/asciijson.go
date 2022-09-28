package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Asciijson(c *gin.Context) {

	data := map[string]interface{}{
		"lang": "GO语言",
		"tag":  "<br>",
	}

    c.AsciiJSON(http.StatusOK, data)
}
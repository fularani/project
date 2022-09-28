package main

import (
    "github.com/gin-gonic/gin"
	"project/utils"
)

func main(){
	r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	r.GET("/ping",utils.Ping)

	r.GET("/someJSON",utils.Asciijson)

	r.GET("/albums",utils.GetAlbums)

	r.POST("/albums", utils.AddAlbums)

	r.GET("/albums/:id", utils.GetAlbumByID)

	r.Run() // listen and serve on 0.0.0.0:8080
}
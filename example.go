package main

import (
    "github.com/gin-gonic/gin"
	"project/utils"
	"project/pkg/handlers"
	"project/pkg/db"
	"log"
)

func main(){
	r := gin.Default()	
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	v1 := r.Group("/v1")
	{
        v1.GET("/ping",utils.Ping)
	    v1.GET("/someJSON",utils.Asciijson)
	    v1.GET("/albums",utils.GetAlbums)
	    v1.POST("/albums", utils.AddAlbums)
	    v1.GET("/albums/:id", utils.GetAlbumByID)
	}

	DB := db.Init()
    h := handlers.New(DB)

	v2 := r.Group("/v2")
	{
        v2.GET("/books",h.GetAllBooks)
	    v2.GET("/books/:id",h.GetBook)
	    v2.POST("/books",h.AddBook)
	    v2.PUT("/books/:id", h.UpdateBook)
	    v2.DELETE("/books/:id", h.DeleteBook)
	}
	
	log.Println("API is running!")
	r.Run() // listen and serve on 0.0.0.0:8080    
}
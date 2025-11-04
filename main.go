package main

import (
	"gin-quickstart/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/albums", controllers.GetAlbums)
	router.POST("/addalbum", controllers.AddAlbum)
	router.GET("/albums/:id", controllers.GetAlbumById)
	router.Run() // listens on 0.0.0.0:8080 by default

}

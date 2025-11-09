package main

import (
	"gin-quickstart/authentication"
	"gin-quickstart/controllers"
	"gin-quickstart/database"
	"gin-quickstart/models"
	"time"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	clerk.SetKey("sk_test_ba5sxOHYqDhFGBruGLxXCpswGK6UdNQEkQf7vF7nKq")
	database.Connect()
	database.DB.AutoMigrate(&models.Album{}, &models.User{})
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(authentication.ClerkAuthMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/albums", controllers.GetAlbums)
	router.POST("/addalbum", controllers.AddAlbum)
	router.GET("/albums/:id", controllers.GetAlbumById)

	router.GET("/users", controllers.GetAllUsers)
	router.POST("/users", controllers.AddUser)

	router.Run() // listens on 0.0.0.0:8080 by default

}

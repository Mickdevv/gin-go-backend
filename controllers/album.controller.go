package controllers

import (
	"gin-quickstart/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Albums)
}

func AddAlbum(c *gin.Context) {
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	models.Albums = append(models.Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range models.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return

		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

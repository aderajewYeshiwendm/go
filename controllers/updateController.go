package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go/initializers"
	"github.com/go/models"
)

func PutPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	initializers.DB.First(&post, id)

	c.BindJSON(&post)
	initializers.DB.Save(&post)

	c.JSON(http.StatusOK, gin.H{"post": post})
}

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go/initializers"
	"github.com/go/models"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func GetPost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	initializers.DB.First(&post, id)

	c.JSON(http.StatusOK, gin.H{"post": post})
}

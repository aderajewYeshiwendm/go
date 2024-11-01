package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go/initializers"
	"github.com/go/models"
)

func DeletePost(c *gin.Context) {
	var post models.Post
	id := c.Param("id")
	initializers.DB.First(&post, id)

	initializers.DB.Delete(&post)

	c.JSON(http.StatusOK, gin.H{"post": post})
}

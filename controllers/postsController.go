package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go/initializers"
	"github.com/go/models"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	posts := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&posts)
	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

}

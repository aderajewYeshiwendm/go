package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go/controllers"
	"github.com/go/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.PUT("/posts/:id", controllers.PutPost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	r.Run()
}

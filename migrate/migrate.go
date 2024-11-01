package main

import (
	"github.com/go/initializers"
	"github.com/go/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	initializers.DB.AutoMigrate(&models.Post{})
}

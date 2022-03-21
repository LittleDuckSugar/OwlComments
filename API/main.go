package main

import (
	"context"
	"owlcomments/controller"
	"owlcomments/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	// Init the db connection
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(repository.UriDB))

	if err != nil {
		panic(err)
	}

	repository.Client = client
}

func main() {
	// Tells to gin if we are in a dev environment or not
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	// Tells to gin to force color in shell
	gin.ForceConsoleColor()

	router := gin.Default()

	router.GET("/target/:targetId/comments", controller.GetComments)
	router.POST("/target/:targetId/comments", controller.PostComment)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	// router.Run()
	router.Run(":3000")
}

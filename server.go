package main

import (
	"github.com/GoGinApi/v2/controller"
	"github.com/GoGinApi/v2/middleware"
	"github.com/GoGinApi/v2/repository"
	"github.com/GoGinApi/v2/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

var (
	videoRepository = repository.NewVideoRepository()
	videoService    = service.New(videoRepository)
	videoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	server := SetupRouter()
	_ = server.Run(":8082")
}

func SetupRouter() *gin.Engine {
	defer videoRepository.CloseDB()

	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger())

	server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "pong",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "video input is valid"})
		}
	})

	server.PUT("/videos/:id", func(ctx *gin.Context) {
		err := videoController.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "video input is valid"})
		}
	})

	server.DELETE("/videos/:id", func(ctx *gin.Context) {
		err := videoController.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "video input is valid"})
		}
	})
	return server
}

package main

import (
	"io"
	"net/http"
	"os"

	"github.com/GoGinApi/v2/controller"
	"github.com/GoGinApi/v2/middleware"
	"github.com/GoGinApi/v2/repository"
	"github.com/GoGinApi/v2/service"
	"github.com/gin-gonic/gin"
)

var (
	userRepository = repository.DatabaseConnection()
	userService    = service.NewUser(userRepository)
	userController = controller.NewUser(userService)

	expenseRepository = repository.DatabaseConnection()
	expenseService    = service.NewExpense(expenseRepository)
	expenseController = controller.NewExpense(expenseService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	server := setupRouter()
	_ = server.Run(":8082")
}

//SetSetupRouter initializing server
func setupRouter() *gin.Engine {
	// defer userRepository.CloseDB()

	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger())
	server.Use(middleware.Cors())
	server.Use(middleware.RequestIDMiddleware())

	v1 := server.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"ping": "pong",
			})
		})

		v1.GET("/users", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, userController.GetAllUsers())
		})

		v1.GET("/expense", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, expenseController.GetAllExpense())
		})

		v1.GET("/users/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, userController.GetUser(ctx))
		})

		v1.POST("/expense", func(ctx *gin.Context) {
			err := expenseController.AddExpense(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "expense input is valid"})
			}
		})

		v1.POST("/users", func(ctx *gin.Context) {
			err := userController.InsertUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "user input is valid"})
			}
		})

		v1.PUT("/users/:id", func(ctx *gin.Context) {
			err := userController.UpdateUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "video input is valid"})
			}
		})

		v1.DELETE("/users/:id", func(ctx *gin.Context) {
			err := userController.DeleteUser(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "video input is valid"})
			}
		})
	}

	return server
}

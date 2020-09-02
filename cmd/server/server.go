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
//nolint:funlen
func setupRouter() *gin.Engine {
	// defer userRepository.CloseDB()

	setupLogOutput()

	server := gin.New()
	server.Use(gin.Recovery(), middleware.Logger())
	//server.Use(middleware.Cors())
	server.Use(middleware.RequestIDMiddleware())

	v1 := server.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"ping": "pong",
			})
		})

		v1.GET("/getExpense", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, expenseController.GetAllExpense())
		})

		v1.POST("/addExpense", func(ctx *gin.Context) {
			err := expenseController.AddExpense(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "expense input is valid"})
			}
		})

		v1.GET("/session", func(ctx *gin.Context) {
			user, isAuthenticated := controller.AuthMiddleware(ctx, []byte("secret"))
			if !isAuthenticated {
				ctx.JSON(http.StatusUnauthorized, gin.H{"success": false, "msg": "unauthorized"})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"success": isAuthenticated, "user": user})
		})

		v1.POST("/login", func(ctx *gin.Context) {
			err := userController.Login(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Please provide a valid password"})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "User Logged in"})
			}
		})

		v1.POST("/register", func(ctx *gin.Context) {
			err := userController.Create(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "User not able to create"})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "New user is created"})
			}
		})

		v1.POST("/createReset", func(ctx *gin.Context) {
			err := userController.InitiatePasswordReset(ctx)
			if err != "" {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not able to reset"})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "user input is valid"})
			}
		})

		v1.POST("/resetPassword", func(ctx *gin.Context) {
			err := userController.ResetPassword(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": 0})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "user input is valid"})
			}
		})
	}
	return server
}

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

	accountRepository = repository.DatabaseConnection()
	accountService    = service.NewAccountService(accountRepository)
	accountController = controller.NewAccount(accountService)
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
	server.Use(middleware.Cors())
	server.Use(middleware.RequestIDMiddleware())

	server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ping": "pong",
		})
	})

	server.POST("/api/login", func(ctx *gin.Context) {
		err := userController.Login(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "User Logged in"})
		}
	})

	v1 := server.Group("/api/v1")

	v1.Use(middleware.AuthMiddleware())
	{
		v1.GET("/getExpense", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, expenseController.GetAllExpense())
		})

		v1.GET("/getExpense/:id", func(ctx *gin.Context) {
			res, err := expenseController.GetExpense(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		v1.PUT("/updateExpense/:id", func(ctx *gin.Context) {
			err := expenseController.UpdateExpense(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "expense details updated"})
			}
		})

		v1.DELETE("/deleteExpense/:id", func(ctx *gin.Context) {
			err := expenseController.DeleteExpense(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "expense deleted"})
			}
		})

		v1.POST("/addExpense", func(ctx *gin.Context) {
			err := expenseController.AddExpense(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "expense added successfully"})
			}
		})

		v1.GET("/getAccountDetails/:id", func(ctx *gin.Context) {
			res, err := accountController.GetAccountDetails(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, res)
			}
		})

		v1.PUT("/updateAccountDetails/:id", func(ctx *gin.Context) {
			err := accountController.UpdateAccountDetails(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "account details updated"})
			}
		})

		v1.POST("/addAccountDetails", func(ctx *gin.Context) {
			err := accountController.AddAccountDetails(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "account details added"})
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

		v1.POST("/logout", func(ctx *gin.Context) {
			userController.Logout(ctx)
			ctx.JSON(http.StatusOK, gin.H{"message": "User Logged out"})
		})

		v1.POST("/register", func(ctx *gin.Context) {
			err := userController.Create(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "New user is created"})
			}
		})

		v1.POST("/createReset", func(ctx *gin.Context) {
			res, err := userController.InitiatePasswordReset(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not able to reset"})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": res})
			}
		})

		v1.POST("/resetPassword/:id", func(ctx *gin.Context) {
			err := userController.ResetPassword(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "user input invalid"})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "password reset request successful"})
			}
		})
	}
	return server
}

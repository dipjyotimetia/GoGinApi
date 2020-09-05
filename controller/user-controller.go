package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/errors"
	"github.com/GoGinApi/v2/pkg/utils"
	"github.com/GoGinApi/v2/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var jwtKey = []byte("secret")

//Claims jwt claims struct
type Claims struct {
	entity.User
	jwt.StandardClaims
}

//UserController having user function
type UserController interface {
	InitiatePasswordReset(ctx *gin.Context) (string, error)
	ResetPassword(ctx *gin.Context) error
	Create(ctx *gin.Context) error
	Login(ctx *gin.Context) error
	CheckUserExist(ctx *gin.Context) bool
	CheckAndRetrieveUserIDViaEmail(ctx *gin.Context) (int, bool)
}

//userController is having service
type userController struct {
	service service.UserService
}

var _ *validator.Validate

//NewUser implementing userController
func NewUser(service service.UserService) UserController {
	_ = validator.New()
	return &userController{service: service}
}

//InitiatePasswordReset email with reset url
func (uc *userController) InitiatePasswordReset(ctx *gin.Context) (string, error) {
	var createReset entity.CreateReset
	ctx.ShouldBindJSON(&createReset)
	if id, ok := uc.service.CheckAndRetrieveUserIDViaEmail(createReset); ok {
		link := fmt.Sprintf("%s/resetPassword/%d", "http://localhost:8082/api/v1", id)
		return link, nil
		//Reset link is returned in json response for testing purposes since no email service is integrated
	}
	return "", fmt.Errorf("please provide valid user pass")
}

//ResetPassword password reset
func (uc *userController) ResetPassword(ctx *gin.Context) error {
	var resetPassword entity.ResetPassword
	ctx.ShouldBindJSON(&resetPassword)
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	resetPassword.ID = int(id)
	return uc.service.ResetPassword(resetPassword)
}

//Create new user
func (uc *userController) Create(ctx *gin.Context) error {
	var user entity.Register
	ctx.ShouldBindJSON(&user)
	exists := uc.CheckUserExist(ctx)

	valErr := utils.ValidateUser(user, errors.ValidationErrors)
	if !exists {
		valErr = append(valErr, "email already exists")
	}
	if len(valErr) > 0 {
		return fmt.Errorf(valErr[0])
	}
	return uc.service.Create(user)
}

// Login controller
func (uc *userController) Login(ctx *gin.Context) error {
	var user entity.Login
	ctx.ShouldBindJSON(&user)
	var name, email, password, createdAt, updatedAt string

	//expiration time of the token ->30 mins
	expirationTime := time.Now().Add(30 * time.Minute)

	// Create the JWT claims, which includes the User struct and expiry time
	claims := &Claims{

		User: entity.User{
			Name: name, Email: email, CreatedAt: createdAt, UpdatedAt: updatedAt,
		},
		StandardClaims: jwt.StandardClaims{
			//expiry time, expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT token string
	tokenString, err := token.SignedString(jwtKey)
	errors.HandleErr(ctx, err)
	// c.SetCookie("token", tokenString, expirationTime, "", "*", true, false)
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	return uc.service.Login(name, email, password, createdAt, updatedAt, user)
}

//CheckUserExist check user exists
func (uc *userController) CheckUserExist(ctx *gin.Context) bool {
	var register entity.Register
	ctx.ShouldBindJSON(&register)
	return uc.service.CheckUserExist(register)
}

//CheckAndRetrieveUserIDViaEmail -1 as ID if the user doesn't exist in the table
func (uc *userController) CheckAndRetrieveUserIDViaEmail(ctx *gin.Context) (int, bool) {
	var createReset entity.CreateReset
	ctx.ShouldBindJSON(&createReset)
	return uc.service.CheckAndRetrieveUserIDViaEmail(createReset)
}

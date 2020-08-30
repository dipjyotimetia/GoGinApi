package controller

//import (
//	"net/http"
//	"strconv"
//
//	"github.com/GoGinApi/v2/entity"
//	"github.com/GoGinApi/v2/service"
//	"github.com/gin-gonic/gin"
//	"github.com/go-playground/validator/v10"
//)
//
////VideoController is
//type VideoController interface {
//	FindAll() []entity.Video
//	Save(ctx *gin.Context) error
//	Update(ctx *gin.Context) error
//	Delete(ctx *gin.Context) error
//	ShowAll(ctx *gin.Context)
//}
//
////Controller is
//type Controller struct {
//	service service.VideoService
//}
//
//var validate *validator.Validate
//
////New is
//func New(service service.VideoService) VideoController {
//	validate = validator.New()
//	return &Controller{
//		service: service,
//	}
//}
//
////FindAll is
//func (c *Controller) FindAll() []entity.Video {
//	return c.service.FindAll()
//}
//
////Save is
//func (c *Controller) Save(ctx *gin.Context) error {
//	var video entity.Video
//	err := ctx.ShouldBindJSON(&video)
//	if err != nil {
//		return err
//	}
//	err = validate.Struct(video)
//	if err != nil {
//		return err
//	}
//	c.service.Save(video)
//	return nil
//}
//
////ShowAll is
//func (c *Controller) ShowAll(ctx *gin.Context) {
//	videos := c.service.FindAll()
//	data := gin.H{
//		"title":  "Video Page",
//		"videos": videos,
//	}
//	ctx.HTML(http.StatusOK, "index.html", data)
//}
//
////Update is
//func (c *Controller) Update(ctx *gin.Context) error {
//	var video entity.Video
//	err := ctx.ShouldBindJSON(&video)
//	if err != nil {
//		return err
//	}
//	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
//	if err != nil {
//		return err
//	}
//	video.ID = id
//	err = validate.Struct(video)
//	if err != nil {
//		return err
//	}
//	c.service.Update(video)
//	return nil
//}
//
////Delete is ...
//func (c Controller) Delete(ctx *gin.Context) error {
//	var video entity.Video
//	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
//	if err != nil {
//		return err
//	}
//	video.ID = id
//	c.service.Delete(video)
//	return nil
//}

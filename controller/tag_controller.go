package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/imvahid/golang-api-sample/data/request"
	"github.com/imvahid/golang-api-sample/data/response"
	"github.com/imvahid/golang-api-sample/helper"
	"github.com/imvahid/golang-api-sample/service"
	"net/http"
	"strconv"
)

type TagController struct {
	tagService service.TagService
}

func NewTagController(tagService service.TagService) *TagController {
	return &TagController{
		tagService: tagService,
	}
}

func (controller *TagController) Create(c *gin.Context) {
	createTagRequest := request.CreateTagRequest{}
	err := c.ShouldBindJSON(&createTagRequest)
	helper.ErrorPanic(err)

	controller.tagService.Create(createTagRequest)

	data := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, data)
}

func (controller *TagController) Update(c *gin.Context) {
	updateTagRequest := request.UpdateTagRequest{}
	err := c.ShouldBindJSON(&updateTagRequest)
	helper.ErrorPanic(err)

	tagID := c.Param("tagID")
	id, err := strconv.Atoi(tagID)
	helper.ErrorPanic(err)
	updateTagRequest.ID = id

	controller.tagService.Update(updateTagRequest)

	data := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, data)
}

func (controller *TagController) Delete(c *gin.Context) {
	tagID := c.Param("tagID")
	id, err := strconv.Atoi(tagID)
	helper.ErrorPanic(err)

	controller.tagService.Delete(id)

	data := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, data)
}

func (controller *TagController) FindByID(c *gin.Context) {
	tagID := c.Param("tagID")
	id, err := strconv.Atoi(tagID)
	helper.ErrorPanic(err)

	result := controller.tagService.FindByID(id)

	data := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, data)
}

func (controller *TagController) FindAll(c *gin.Context) {
	result := controller.tagService.FindAll()

	data := response.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, data)
}

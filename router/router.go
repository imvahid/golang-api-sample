package router

import (
	"github.com/gin-gonic/gin"
	"github.com/imvahid/golang-api-sample/controller"
	"net/http"
)

func NewRouter(tagController *controller.TagController) *gin.Engine {
	router := gin.Default()

	router.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome Home!")
	})

	// Base router
	baseRouter := router.Group("/api")

	// Tags router
	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.POST("", tagController.Create)
	tagsRouter.PATCH("/:tagID", tagController.Update)
	tagsRouter.DELETE("/:tagID", tagController.Delete)
	tagsRouter.GET("/:tagID", tagController.FindByID)
	tagsRouter.GET("", tagController.FindAll)

	return router
}

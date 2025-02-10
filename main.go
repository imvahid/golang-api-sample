package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imvahid/golang-api-sample/helper"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	log.Info().Msg("Started Server!")
	routes := gin.Default()

	routes.GET("", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome Home!")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}

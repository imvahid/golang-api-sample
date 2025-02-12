package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/imvahid/golang-api-sample/config"
	"github.com/imvahid/golang-api-sample/controller"
	"github.com/imvahid/golang-api-sample/helper"
	"github.com/imvahid/golang-api-sample/model"
	"github.com/imvahid/golang-api-sample/repository"
	"github.com/imvahid/golang-api-sample/router"
	"github.com/imvahid/golang-api-sample/service"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()
	helper.ErrorPanic(err)
}

func main() {
	log.Info().Msg("Started Server!")

	// Set mode
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = gin.ReleaseMode // Default to release mode if not set
	}
	gin.SetMode(ginMode)

	// Database
	db := config.MySQLConnection()
	validate := validator.New()
	errMigrate := db.Table("tags").AutoMigrate(&model.Tag{})
	helper.ErrorPanic(errMigrate)

	// Repository
	tagRepository := repository.NewTagRepositoryImplementation(db)

	// Service
	tagService := service.NewTagServiceImplementation(tagRepository, validate)

	// Controller
	tagController := controller.NewTagController(tagService)

	// Router
	routes := router.NewRouter(tagController)

	server := &http.Server{
		Addr:    ":" + os.Getenv("APP_PORT"),
		Handler: routes,
	}

	errServe := server.ListenAndServe()
	helper.ErrorPanic(errServe)
}

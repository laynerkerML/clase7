package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/laynerkerML/clase7/cmd/service/handler"
	"github.com/laynerkerML/clase7/docs"
	"github.com/laynerkerML/clase7/internal/users"
	"github.com/laynerkerML/clase7/pkg/store"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This Api Handle MELI Product
// @termsOfService hola

// @contact.name API Support
// @contact.url https://hola.com

// @license.name Apache 2.0
// @license.url https://github.com/

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := store.New(store.FileType, "users.json")

	router := gin.Default()

	repository := users.NewRepository(db)
	service := users.NewService(repository)
	userHandler := handler.NewUser(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRouter := router.Group("api/v1/users")
	{
		userRouter.GET("/", userHandler.ValidationToken(), userHandler.GetAll())
		userRouter.POST("/", userHandler.ValidationToken(), userHandler.Save())
		userRouter.PUT("/:id", userHandler.ValidationToken(), userHandler.Update())
		userRouter.PATCH("/:id", userHandler.ValidationToken(), userHandler.Patch())
		userRouter.DELETE("/:id", userHandler.ValidationToken(), userHandler.Delete())
	}
	router.Run()
}

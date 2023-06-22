package main

import (
	docs "crud/docs"
	"crud/modules/user"
	"crud/utils/db"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

// @title API Documentation
// @Version 1.0.0
// @description Swagger API for Golang Project Blueprint.
// @contact.name API Support
// @contact.email bowo_bp@outlook.com
// @host localhost:8081
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	router := gin.New()

	// open connection db
	dbCrud := db.GormMysql(os.Getenv("CRUD_DSN"))

	fmt.Println("database connected..!")
	docs.SwaggerInfo.BasePath = "/api/v1"
	versionRoute := router.Group("api/v1")

	userHandler := user.NewRouter(dbCrud)
	userHandler.Handle(versionRoute)

	// Server Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	errRouter := router.Run(os.Getenv("PORT"))
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}

}

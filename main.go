package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cdeor/url-shortener/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/genproto/googleapis/maps/routes/v1"
)

func main() {

	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	router := gin.Default()

	setupRouters(router)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(router.Run(":" + port))

}

func setupRouters(router *gin.Engine) {
	router.POST("/api/v1", routes.ShortenURL)
	router.GET("/api/v1/:shortID", routes.GetByShortID)
}

package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

func rootHandler(context *gin.Context) {
	currentTime := time.Now()
	currentTime.Format("20060102150405")
	context.JSON(200, gin.H{
		"currentTime": currentTime,
		"text":        "hello world",
	})
}

func main() {
	port := os.Getenv("PORT")
	router := gin.Default()
	if port == "" {
		port = "8000"
	}
	router.GET("/", rootHandler)
	router.Run(":" + port)
}

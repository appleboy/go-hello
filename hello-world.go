package main

import (
	"github.com/gin-gonic/gin"
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
	router := gin.Default()
	router.GET("/", rootHandler)
	router.Run(":8000")
}

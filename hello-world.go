package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func rootHandler(context *gin.Context) {
	current_time := time.Now()
	current_time.Format("20060102150405")
	context.JSON(200, gin.H{
		"current_time": current_time,
	})
}

func main() {
	router := gin.Default()
	router.GET("/", rootHandler)
	router.Run(":8000")
}

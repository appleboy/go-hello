package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"time"
)

func rootHandler(context *gin.Context) {
	currentTime := time.Now()
	currentTime.Format("20060102150405")
	context.JSON(200, gin.H{
		"current_time": currentTime,
		"text":         "Hello World",
	})
}

func GetMainEngine() *gin.Engine {
	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/", rootHandler)

	return r
}

func RunHTTPServer() error {
	port := flag.String("port", "8000", "The port for the mock server to listen to")

	// Parse all flag
	flag.Parse()

	err := GetMainEngine().Run(":" + *port)

	return err
}

func main() {
	RunHTTPServer()
}

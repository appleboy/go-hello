package main

import (
	"flag"
	"time"

	"github.com/gin-gonic/gin"
)

func rootHandler(c *gin.Context) {
	currentTime := time.Now()
	currentTime.Format("20060102150405")
	c.JSON(200, gin.H{
		"current_time": currentTime,
		"text":         "Hello World",
	})
}

// GetMainEngine is default router engine using gin framework.
func GetMainEngine() *gin.Engine {
	r := gin.New()

	r.GET("/", rootHandler)

	return r
}

// RunHTTPServer List 8000 default port.
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

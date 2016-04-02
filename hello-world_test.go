package main

import (
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"
	"time"
	"io/ioutil"
)

func testRequest(t *testing.T, url string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	assert.NoError(t, err)

	_, ioerr := ioutil.ReadAll(resp.Body)
	assert.NoError(t, ioerr)
	assert.Equal(t, "200 OK", resp.Status, "should get a 200")
}

func TestGinHelloWorld(t *testing.T) {
	r := gofight.New()

	r.GET("/").
		Run(GetMainEngine(), func(r gofight.HttpResponse, rq gofight.HttpRequest) {
			data := []byte(r.Body.String())

			value, _ := jsonparser.GetString(data, "text")

			assert.Equal(t, "Hello World", value)
			assert.Equal(t, http.StatusOK, r.Code)
		})
}

func TestRunNormalServer(t *testing.T) {
	router := gin.New()

	go func() {
		assert.NoError(t, RunHTTPServer())
	}()
	// have to wait for the goroutine to start and run the server
	// otherwise the main thread will complete
	time.Sleep(5 * time.Millisecond)

	assert.Error(t, router.Run(":8000"))
	testRequest(t, "http://localhost:8000/")
}

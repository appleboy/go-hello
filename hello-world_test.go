package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	type Message struct {
		Text        string `json:"text"`
		CurrentTime string `json:"current_time"`
	}

	gin.SetMode(gin.TestMode)
	r := GetMainEngine()
	req, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// pase json body.
	dec := json.NewDecoder(strings.NewReader(w.Body.String()))
	var exp Message
	if err := dec.Decode(&exp); err != nil {
		log.Fatal(err)
	}

	if exp.Text != "hello world" {
		t.Errorf("JSON Text value didn't return hello world")
	}

	if w.Code != http.StatusOK {
		fmt.Printf("%d - %s", w.Code, w.Body.String())
		t.Errorf("Home page didn't return %v", http.StatusOK)
	}
}

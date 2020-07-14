package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Start(endPoint string) http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "elastic"})
	})
	return e
}

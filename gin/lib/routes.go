package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hcHandler(c *gin.Context) {
	workingMessage := map[string]string{
		"status": "WORKING",
	}
	c.JSON(http.StatusOK, workingMessage)
}

func makeRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/", hcHandler)

	return r
}

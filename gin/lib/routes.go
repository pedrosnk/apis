package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type hcResponse struct {
	Status string `json:"status"`
}

func hcHandler(c *gin.Context) {
	workingMessage := hcResponse{Status: "WORKING"}
	c.JSON(http.StatusOK, workingMessage)
}

func makeRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/healthcheck", hcHandler)

	return r
}

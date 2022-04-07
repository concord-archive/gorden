package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.SetTrustedProxies(nil)
	server.GET("/__development/ping", ping)

	log.Fatal(server.Run(":5000"))
}

type DefaultMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"cookie": "ping"})
}

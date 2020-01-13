package main

import (
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var shPath = os.Args[1]

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://192.168.0.16:8080"}
	config.AllowMethods = []string{"GET"}
	router.Use(cors.New(config))

	api := router.Group("/api/v1")
	{
		api.GET("/heyson", heySon)
	}
	router.Run(":3000")
}

func heySon(c *gin.Context) {
	err := exec.Command(shPath).Start()
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		c.Status(http.StatusOK)
	}
}

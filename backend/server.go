package main

import (
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

var shPath = os.Args[1]

func main() {
	router := gin.Default()
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

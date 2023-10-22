package main

import (
	//"container/vector"
	"net/http"
	//"os"

	"github.com/gin-gonic/gin"
)

type SubmissionTemplate struct {
	Name             string
	ValidExtension   string
	InvalidExtension string
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Index",
		})
	})

	r.GET("/template", func(c *gin.Context) {
		//filename := c.Query("filename")
		//fileout, err := os.Open(filename + ".json")

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

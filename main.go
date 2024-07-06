package main

import (
	"net/http"

	"github.com/KevinMi2023p/go_simple_api/jsonProcessing"
	"github.com/gin-gonic/gin"
)

func main() {
	// Your code here
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/receipts/process", func(c *gin.Context) {
		var requestBody map[string]interface{}
		err := c.ShouldBindJSON(&requestBody)
		if err != nil {
			//c.(http.StatusBadRequest, The receipt is invalid)
			c.String(http.StatusBadRequest, "The receipt is invalid")
			return
		}

		err = jsonProcessing.CheckJSONStructure(requestBody)
		if err != nil {
			c.String(http.StatusBadRequest, "The receipt is invalid")
			return
		}

		c.JSON(http.StatusOK, requestBody)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

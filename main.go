package main

import (
	"net/http"

	"github.com/KevinMi2023p/go_simple_api/jsonProcessing"
	"github.com/gin-gonic/gin"
)

var mockDatabase = make(map[string]int)

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

		generatedID := jsonProcessing.HashReciet(requestBody)
		mockDatabase[generatedID] = jsonProcessing.GetScore()

		output := map[string]interface{}{
			"id": generatedID,
		}

		c.JSON(http.StatusOK, output)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

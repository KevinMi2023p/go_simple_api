package main

import (
	"net/http"

	"github.com/KevinMi2023p/go_simple_api/jsonProcessing"
	"github.com/gin-gonic/gin"
)

var mockDatabase = make(map[string]int)

func main() {
	// Your code here
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/receipts/process", func(c *gin.Context) {
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

	router.GET("/receipts/:id/points", func(c *gin.Context) {
		id := c.Param("id")
		points, ok := mockDatabase[id]
		if !ok {
			c.String(http.StatusNotFound, "No receipt found for that id")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"points": points,
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

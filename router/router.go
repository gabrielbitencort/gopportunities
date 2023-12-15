package router

import "github.com/gin-gonic/gin"

func Initialize() {

	// Initialize Router using default settings
	router := gin.Default()

	// Define Route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Run API
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}

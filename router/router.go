package router

import "github.com/gin-gonic/gin"

func Initialize() {

	// Initialize Router using default settings
	router := gin.Default()

	// Initialize Route
	initializeRoutes(router)

	// Run server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}

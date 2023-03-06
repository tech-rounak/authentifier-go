package main

import (
	"os"

	routes "github.com/tech-rounak/jwt-authentication/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "success true for api2"})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "success true for api2"})
	})

	router.Run(":" + port)
}

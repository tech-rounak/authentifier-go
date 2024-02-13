package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/tech-rounak/jwt-authentication/controllers"
	"github.com/tech-rounak/jwt-authentication/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/upload", controller.UploadImage())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/user/:id", controller.GetUser())
}

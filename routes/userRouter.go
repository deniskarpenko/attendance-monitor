package routes

import (
	"github.com/deniskarpenko/attendance-monitor/controllers"
	"github.com/deniskarpenko/attendance-monitor/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("users/:user_id", controllers.GetUser())
}

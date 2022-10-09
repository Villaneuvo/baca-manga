package routes

import (
	"baca-manga/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.RouterGroup) {
	router.GET("/ping", controllers.Ping)
	router.GET("/home", controllers.Home)
}

package router

import (
	"app-log/app/http/controller"

	"github.com/gin-gonic/gin"
)

func LoadLog(e *gin.Engine) {
	api := e.Group("api")
	{
		log := api.Group("log")
		{
			log.GET("/index", controller.GetUser)
		}
	}
}

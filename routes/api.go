package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadTest(e *gin.Engine) {
	// e.GET("/ping", controller.GetUser)
	// e.POST("/user", controller.RegisterUser)
	// api := e.Group("api")
	// {
	// 	api.POST("/login", controller.Login)

	// 	// 用户相关
	// 	user := api.Group("/user")
	// 	{
	// 		user.GET("/info", controller.Info)
	// 	}

	// 	buildings := api.Group("/buildings")
	// 	{
	// 		buildings.GET("/list", controller.DictBuildingsList)
	// 	}

	e.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})
}

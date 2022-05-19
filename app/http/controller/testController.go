package controller

import "github.com/gin-gonic/gin"

func GetUser(c *gin.Context) {
	username := c.Query("name")
	age := c.DefaultQuery("age", "11")

	// data := make(map[string]interface{})
	data := make(map[string]string)
	data["username"] = username
	data["age"] = age

	c.JSON(200, gin.H{
		"message": "ok",
		"data":    data,
	})
}

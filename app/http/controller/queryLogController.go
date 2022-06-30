package controller

import (
	"app-log/app/service"
	mongoDB "app-log/pkg/database/mongoDb"

	"github.com/gin-gonic/gin"
)

func FindOne(c *gin.Context) {

	// name := "test_mongo"
	// table := "my_log" // 表名

	name := "Test"
	table := "test" // 表名
	db, _ := mongoDB.ConnectToDB(name)
	collection := db.Collection(table)
	initRepo := mongoDB.NewMongoRepository(collection)
	initSvc := service.NewMongoSvc(initRepo)

	// 调用
	// initSvc.Find()
	initSvc.Count()
	c.JSON(200, gin.H{"msg": "ok", "status": 0, "data": 1})
	// gin.j
	// initSvc.List()
}

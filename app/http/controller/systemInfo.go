package controller

import (
	"app-log/app/service"
	mongoDB "app-log/pkg/database/mongoDb"
	tools "app-log/pkg/tools/json"

	"github.com/gin-gonic/gin"
)

func GEtSystemInfo(c *gin.Context) {

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
	total := initSvc.Count()
	data := make(map[string]interface{})
	data["total"] = total
	// all document num
	// all db info
	// all collection info
	c.JSON(200, gin.H{"msg": "ok", "status": tools.SUCCESS, "data": 1})
}

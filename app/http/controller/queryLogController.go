package controller

import (
	"app-log/app/service"
	mongoDB "app-log/pkg/database/mongoDb"
	tools "app-log/pkg/tools/json"
	"context"
	"fmt"

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
	allNames, _ := db.Client().ListDatabaseNames(context.Background(), struct{}{})
	fmt.Println(allNames)

	c.JSON(200, gin.H{"msg": "ok", "status": tools.SUCCESS, "data": 1})
	// gin.j
	// initSvc.List()
}

func GetAllDbNames(c *gin.Context) {

	name := "Test"
	db, _ := mongoDB.ConnectToDB(name)

	allNames, _ := db.Client().ListDatabaseNames(context.Background(), struct{}{})

	c.JSON(200, gin.H{"msg": "ok", "status": tools.SUCCESS, "data": allNames})
}

func GetAllClsName(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "ok", "status": tools.SUCCESS, "data": service.GetAllClsName()})
}

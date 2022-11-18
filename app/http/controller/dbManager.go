package controller

import (
	"app-log/app/service"
	tools "app-log/pkg/tools/json"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func GetDbList(c *gin.Context) {
	GetConn()
	dbs, err := conn.Client.ListDatabaseNames(c, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	type list = map[string]interface{}
	var listArr []list

	var document bson.M
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	for _, v := range dbs {
		conn.Client.Database(v).RunCommand(
			ctx,
			bsonx.Doc{{"dbStats", bsonx.Int32(1)}},
		).Decode(&document)

		var l = make(list)

		l["db"] = v
		l["clsNum"] = document["collections"]
		l["size"] = document["dataSize"]
		if v == "admin" || v == "config" || v == "local" {
			l["disabled"] = true
		} else {
			l["disabled"] = false
		}

		listArr = append(listArr, l)
	}
	c.JSON(200, gin.H{"msg": "ok", "status": tools.SUCCESS, "data": listArr})
}

func CreateDB(c *gin.Context) {
	// 定义接收数据的结构体
	type DbInfo struct {
		// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
		DbName  string `form:"dbName" json:"dbName" uri:"dbName" xml:"dbName" binding:"required"`
		ClsName string `form:"clsName" json:"clsName" uri:"clsName" xml:"clsName" binding:"required"`
	}

	// 声明接收的变量
	var form DbInfo

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	GetConn()
	opts := options.CreateCollection().SetCapped(true).SetSizeInBytes(1024)
	err := conn.Client.Database(form.DbName).CreateCollection(c, form.ClsName, opts)
	if err != nil {
		c.JSON(200, gin.H{"msg": "ok", "status": tools.ERROR, "data": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "ok", "status": tools.SUCCESS, "data": nil})
}

func DeleteDB(c *gin.Context) {
	// 定义接收数据的结构体
	/* type DbInfo struct {
		// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
		DbName string `form:"dbName" json:"dbName" uri:"dbName" xml:"dbName" binding:"required"`
	}

	// 声明接收的变量
	var form DbInfo

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	GetConn()
	err := conn.Client.Database(form.DbName).Drop(c)
	if err != nil {
		c.JSON(200, gin.H{"msg": "ok", "status": tools.ERROR, "data": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "ok", "status": tools.SUCCESS, "data": nil}) */
	service.SyncLog("/logs")
}

func TestSync() {
	service.SyncLog("/logs")
}

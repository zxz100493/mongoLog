package controller

import (
	mongoDb "app-log/pkg/database/mongoDb"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func TestC(c *gin.Context) {
	uri := "mongodb://root:123456@127.0.0.1:27017/admin"
	name := "Test"
	maxTime := time.Duration(2) // 链接超时时间
	table := "test"             // 表名
	type Test struct {
		Id    string `bson:"_id"`
		Name  string `bson:"name"`
		Level int    `bson:"level"`
	}

	db, err := mongoDb.ConnectToDB(uri, name, maxTime)
	collection = db.Collection(table)

	var initMongo = new(mongoDb.Mongo)
	initMongo.Collection = collection

	initMongo.Model = Test{
		Id:    "1",
		Name:  "zngw",
		Level: 55,
	}

	if err != nil {
		log.Println("链接数据库有误!")
	} else {
		log.Println("链接成功!")
	}

	mongoDb.AddOne(initMongo)
	mongoDb.Count(initMongo)
	mongoDb.GetList(bson.M{"level": 55}, initMongo)
}

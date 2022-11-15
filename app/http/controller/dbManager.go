package controller

import (
	tools "app-log/pkg/tools/json"
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

		listArr = append(listArr, l)
	}
	c.JSON(200, gin.H{"msg": "ok", "status": tools.SUCCESS, "data": listArr})
}

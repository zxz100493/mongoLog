package controller

import (
	mongoDB "app-log/pkg/database/mongoDb"
	"fmt"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoObject struct {
	Client *mongo.Client
}

var (
	once sync.Once
	conn *mongoObject
)

func NewMongoObject() *mongoObject {
	fmt.Println("NewMongoObject")
	return &mongoObject{
		Client: mongoDB.MongoClicent(),
	}
}

func GetConn() {
	if conn == nil {
		once.Do(func() {
			conn = NewMongoObject()
		})
	}
}

func GetSystemInfo(c *gin.Context) {
	GetConn()
	dbs, err := conn.Client.ListDatabaseNames(c, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", dbs)
}

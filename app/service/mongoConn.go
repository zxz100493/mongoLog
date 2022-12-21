package service

import (
	mongoDB "app-log/pkg/database/mongoDb"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoObject struct {
	Client *mongo.Client
}

var (
	once sync.Once
	Conn *mongoObject
)

func NewMongoObject() *mongoObject {
	fmt.Println("NewMongoObject")
	return &mongoObject{
		Client: mongoDB.MongoClicent(),
	}
}

func GetConn() {
	if Conn == nil {
		once.Do(func() {
			Conn = NewMongoObject()
		})
	}
}

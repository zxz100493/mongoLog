package service

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllClsName() []string {
	dbs, err := Conn.Client.ListDatabaseNames(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var cls []string
	for _, v := range dbs {
		names, err := Conn.Client.Database(v).ListCollectionNames(context.Background(), bson.M{})
		if err != nil {
			log.Println(err)
		}
		for _, nv := range names {
			cls = append(cls, nv)
		}
	}
	return cls
}

func GetAllLogName() {
	// get all log name
	dbs, err := Conn.Client.ListDatabaseNames(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range dbs {
		names, err := Conn.Client.Database(v).ListCollectionNames(context.Background(), bson.M{})
		fmt.Println(names)
		fmt.Println(err)
	}
}

func QueryLog(c *gin.Context) []bson.M {
	name := c.Query("uniqueMark")
	dbName := string(name[0])
	fmt.Println(name)
	fmt.Println(dbName)
	opts := options.Find().SetSort(bson.D{{"datetime", -1}})

	cursor, err := Conn.Client.Database(dbName).Collection(name).Find(c, bson.D{}, opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(err)
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	return results
	// for _, result := range results {
	// 	fmt.Println(result)
	// }
	// return "ok"
}
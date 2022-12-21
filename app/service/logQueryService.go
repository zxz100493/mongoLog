package service

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAllClsName() {
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

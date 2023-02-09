package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

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

func QueryLog(c *gin.Context) map[string]interface{} {
	name := c.Query("uniqueMark")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")
	keyword := c.DefaultQuery("keywords", "")
	startDate := c.DefaultQuery("timeStart", "")
	endDate := c.DefaultQuery("timeEnd", "")
	fmt.Println(startDate, endDate)
	intPage, err := strconv.Atoi(page)
	if err != nil {
		fmt.Println(err)
	}

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		fmt.Println(err)
	}

	skip := int64(intPage*intLimit - intLimit)
	dbName := string(name[0])
	fmt.Println(name)
	fmt.Println(dbName)
	opts := options.Find().SetSort(bson.D{{"datetime", -1}}).SetLimit(int64(intLimit)).SetSkip(skip)
	countOpts := new(options.CountOptions)
	// .FindOptions{Limit: &intLimit, Skip: &skip}
	var total int64
	total, err = Conn.Client.Database(dbName).Collection(name).CountDocuments(c, bson.D{}, countOpts)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.D{}
	fmt.Println("keyword", keyword)
	if keyword != "" {
		newKeyword := strings.Split(keyword, ":")

		// filter = append(filter, bson.E{"context", "/" + keyword + "/"})
		/* filter = append(filter,
		bson.E{
			Key: "context",
			// Value: bson.M{"$regex": primitive.Regex{Pattern: "/a/", Options: "im"}},
			Value: bson.M{"$regex": primitive.Regex{Pattern: "/" + keyword + "/", Options: "im"}},
		}) */
		filter = append(filter,
			bson.E{
				Key: newKeyword[0],
				// Value: bson.M{"$regex": primitive.Regex{Pattern: "/a/", Options: "im"}},
				Value: newKeyword[1],
			})

	}
	if endDate != "" && startDate != "" {
		filter = append(filter, bson.E{Key: "datetime", Value: bson.M{"$gt": startDate, "$lt": endDate}})
	}

	fmt.Println("filter:", filter)

	cursor, err := Conn.Client.Database(dbName).Collection(name).Find(c, filter, opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(err)
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	ret := make(map[string]interface{})
	ret["list"] = results
	ret["total"] = total
	return ret
	// for _, result := range results {
	// 	fmt.Println(result)
	// }
	// return "ok"
}

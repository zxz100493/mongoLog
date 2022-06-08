package controller

import (
	"app-log/app/model"
	mongoDb "app-log/pkg/database/mongoDb"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type LogStruct struct {
	Datetime     string `json:"datetime"`
	Timestamp    string `json:"timestamp"`
	UniqueRemark string `json:"unique_remark"`
	CnRemark     string `json:"cn_remark"`
	Project      string `json:"project"`
	UserId       string `json:"user_id"`
	Path         string `json:"path"`
	Module       string `json:"module"`
	Host         string `json:"host"`
	Url          string `json:"url"`
	Level        string `json:"level"`
	Context      string `json:"context"`
	Backtrace    string `json:"backtrace"`
	PostData     string `json:"postData"`
	GetData      string `json:"getData"`
}

func TestC(c *gin.Context) {
	// log.Println(&config.Instance.Mysql.User)
	name := "Test"
	table := "test" // 表名
	type Test model.Test

	db, err := mongoDb.ConnectToDB(name)
	collection := db.Collection(table)
	type NewStruct = mongoDb.Mongo

	var initMongo = new(NewStruct)
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

	// mongoDb.AddOne(initMongo)
	mongoDb.Count(initMongo)
	mongoDb.GetList(bson.M{"level": 55}, initMongo)

	for _, file := range ScanDir() {
		fmt.Println(file)
	}
}

// scan the directort
func ScanDir() (files []string) {
	dir := "/home/zxz/dnmp/www/rrzuji/console/logV3"
	err := filepath.Walk(dir, visit(&files))
	if err != nil {
		panic(err)
	}
	return files
	// for _, file := range files {
	// fmt.Println(file)
	// }
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		s, err := os.Stat(path)
		if err != nil {
			return nil
		}

		if !s.IsDir() {
			*files = append(*files, path)
		}
		return nil
	}
}

package service

import (
	mongoDB "app-log/pkg/database/mongoDb"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	LogContent struct {
		Datetime       string      `json:"datetime" bson:"datetime"`
		Timestamp      int         `json:"timestamp" bason:"timestamp"`
		NginxRequestId string      `json:"nginx_request_id" bson:"nginx_request_id"`
		UniqueRemark   string      `json:"unique_remark" bson:"unique_remark"`
		CnRemark       string      `json:"cn_remark" bson:"cn_remark"`
		UserId         interface{} `json:"user_id" bson:"user_id"`
		Project        string      `json:"project" bson:"project"`
		Types          string      `json:"type" bson:"type"`
		Path           string      `json:"path" bson:"path"`
		Module         string      `json:"module" bson:"module"`
		Host           string      `json:"host" bson:"host"`
		Url            string      `json:"url" bson:"url"`
		Level          string      `json:"level" bson:"level"`
		Context        interface{} `json:"context" bson:"context"`
		Backtrace      interface{} `json:"backtrace" bson:"backtrace"`
		PostData       interface{} `json:"postData" bson:"postData"`
		GetData        interface{} `json:"getData" bson:"getData"`
	}
	lcAndPath []map[string]LogContent
	tmpMap    map[string]interface{}
)

var (
	ch   = make(chan tmpMap, 20)
	tick <-chan time.Time
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

// func init(){
// 	GetConn()
// }

func GetConn() {
	if conn == nil {
		once.Do(func() {
			conn = NewMongoObject()
		})
	}
}

func SyncLog(dir string) {
	go walkDir(dir)
	go watchChannel()
}

func watchChannel() {
	// batch data
	batchData := []tmpMap{}
loop:
	for {
		select {
		case c, ok := <-ch:
			if !ok {
				break loop
			}
			batchData = append(batchData, c)
			if len(batchData) >= 20 {
				writeDb(batchData)     // insert db
				batchData = []tmpMap{} //clear all
			}
		case <-tick:
			fmt.Println("it is time")
			break loop
		}
	}

}

func walkDir(dir string) {
	for _, entry := range dirents(dir) {
		path := filepath.Join(dir, entry.Name())
		if entry.IsDir() {
			SyncLog(path)
		} else {
			readDir(path)
		}
	}
}
func readDir(path string) {
	fmt.Println("path:", path)
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println("error opening json file")
		return
	}
	defer jsonFile.Close()
	read := bufio.NewReader(jsonFile)

	num := 0
	for {
		line, err := read.ReadString('\n')
		num++
		if err == io.EOF {
			break
		}
		str := []byte(line)
		if err != nil {
			fmt.Println("error reading json file")
			return
		}
		var lc LogContent
		err = json.Unmarshal(str, &lc)
		if err != nil {
			fmt.Println("err--------")
			fmt.Printf("%s", str)
			fmt.Println("line--------")
			fmt.Printf("%s", line)
			log.Fatal("err:", err)
		}
		tmpMap := make(tmpMap)
		tmpMap["path"] = path
		tmpMap["content"] = lc

		ch <- tmpMap

		// fmt.Printf("%s", lc)
		// put to channel and send to client
	}
	// fmt.Println("all_line_is:", num)
}

// write db
func writeDb(data []tmpMap) {
	fmt.Println("xxxxxxxxxxxxxxxxxxxx")

	// divide db
	// var info = make(map[string]map[string][]tmpMap)
	var info = make(map[string]interface{})
	var content = make(map[string][]interface{})
	GetConn()
	// var insertInfo []info
	// i := make(info)
	for _, v := range data {
		path := insertDbName(v["path"])
		dbName := strings.ToLower(string(path[3]))
		clsName := strings.ToLower(path[3:])
		_, ok := content[clsName]
		if !ok {
			content = make(map[string][]interface{})
		}
		content[clsName] = append(content[clsName], v["content"])
		info[dbName] = content
	}
	// insertDb
	for dbKey, v := range info {
		vv := v.(map[string][]interface{})
		for cls, v3 := range vv {
			res, err := conn.Client.Database(dbKey).Collection(cls).InsertMany(context.TODO(), v3)
			if err != nil {
				log.Fatal(err)
			}
			// log.Printf("%v", v3)
			// log.Printf("%v", res)
			fmt.Printf("inserted documents with IDs %v\n", res.InsertedIDs)
			log.Fatal()
		}
	}
}

func insertDbName(pathName interface{}) string {
	path := pathName.(string)
	index := strings.LastIndex(path, "/")
	name := path[index+1:]
	suffix := strings.Split(name, ".")[1]
	newPath := strings.Replace(name, "."+suffix, " ", 1)
	return newPath
}

func dirents(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("read dir:%v\n", err)
		return nil
	}
	return entries
}

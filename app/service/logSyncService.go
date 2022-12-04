package service

import (
	mongoDB "app-log/pkg/database/mongoDb"
	"bufio"
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
		Datetime       string      `json:"datetime"`
		Timestamp      int         `json:"timestamp"`
		NginxRequestId string      `json:"nginx_request_id"`
		UniqueRemark   string      `json:"unique_remark"`
		CnRemark       string      `json:"cn_remark"`
		UserId         interface{} `json:"user_id"`
		Project        string      `json:"project"`
		Types          string      `json:"type"`
		Path           string      `json:"path"`
		Module         string      `json:"module"`
		Host           string      `json:"host"`
		Url            string      `json:"url"`
		Level          string      `json:"level"`
		Context        interface{} `json:"context"`
		Backtrace      interface{} `json:"backtrace"`
		PostData       interface{} `json:"postData"`
		GetData        interface{} `json:"getData"`
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
	type info map[string]interface{}
	// var insertInfo []info
	// i := make(info)
	for _, v := range data {
		path := insertDbName(v["path"])
		dbName := strings.ToLower(string(path[3]))
		// newInfo := i{
		// 	"db" : dbName,
		// 	"collection" : path,
		// 	"data" : v["content"],
		// }
		// insertInfo = append(insertInfo, newInfo)
		// log.Fatal(dbName)
	}

	// insertDb
	// conn.Client.Database(dbName).Collection(path).InsertMany(context.TODO(),v["content"])
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

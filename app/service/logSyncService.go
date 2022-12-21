package service

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		// var lc LogContent
		// var lc bson.Raw
		// var lc []byte
		// var lc bson.D
		// var lc []interface{}

		// err = json.Unmarshal(str, &lc)
		var data bson.D
		err = bson.UnmarshalExtJSON(str, true, &data)
		if err != nil {
			fmt.Println(err)
		}
		if err != nil {
			fmt.Println("err--------")
			fmt.Printf("%s", str)
			fmt.Println("line--------")
			log.Fatal("err1-unmarshal:", err)
		}
		tmpMap := make(tmpMap)
		tmpMap["path"] = path
		tmpMap["content"] = data

		ch <- tmpMap
	}
}

// write db
func writeDb(data []tmpMap) {
	fmt.Println("xxxxxxxxxxxxxxxxxxxx")

	// divide db
	var info = make(map[string]interface{})
	var content = make(map[string][]interface{})
	GetConn()
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
			res, err := Conn.Client.Database(dbKey).Collection(cls).InsertMany(context.TODO(), v3)
			if err != nil {
				log.Fatal(err)
			}
			opts := options.FindOne().SetSort(bson.D{{"_id", 1}})
			var result bson.M
			for _, id := range res.InsertedIDs {
				fmt.Printf("inserted documents with IDs %v\n", id)
				err := Conn.Client.Database(dbKey).Collection(cls).FindOne(context.TODO(),
					bson.D{{"_id", id}},
					opts).Decode(&result)
				if err != nil {
					// ErrNoDocuments means that the filter did not match any documents in
					// the collection.
					if err == mongo.ErrNoDocuments {
						return
					}
					log.Fatal(err)
				}
				// fmt.Printf("found document %v", result)
				// log.Fatal()
			}
		}
	}
}

func insertDbName(pathName interface{}) string {
	path := pathName.(string)
	index := strings.LastIndex(path, "/")
	name := path[index+1:]
	suffix := strings.Split(name, ".")[1]
	clearQuote := strings.Replace(name, ".", "_", strings.Count(name, "."))
	newPath := strings.Replace(clearQuote, "_"+suffix, " ", 1)
	log.Println(path)
	log.Println(strings.TrimSpace(newPath))
	log.Println("----------------")

	return strings.TrimSpace(newPath)
}

func dirents(dir string) []os.DirEntry {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("read dir:%v\n", err)
		return nil
	}
	return entries
}

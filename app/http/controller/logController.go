package controller

import (
	"app-log/app/model"
	mongoDb "app-log/pkg/database/mongoDb"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// type LogStruct struct {
// 	Datetime     string      `json:"datetime"`
// 	Timestamp    int32       `json:"timestamp"`
// 	UniqueRemark string      `json:"unique_remark"`
// 	CnRemark     string      `json:"cn_remark"`
// 	Project      string      `json:"project"`
// 	UserId       int32       `json:"user_id"`
// 	Path         string      `json:"path"`
// 	Module       string      `json:"module"`
// 	Host         string      `json:"host"`
// 	Url          string      `json:"url"`
// 	Level        string      `json:"level"`
// 	Context      interface{} `json:"context"`
// 	Backtrace    interface{} `json:"backtrace"`
// 	PostData     interface{} `json:"postData"`
// 	GetData      interface{} `json:"getData"`
// }
type LogStruct = model.LogStruct

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

	// initMongo.Model = Test{
	// 	Id:    "1",
	// 	Name:  "zngw",
	// 	Level: 55,
	// }

	if err != nil {
		log.Println("链接数据库有误!")
	} else {
		log.Println("链接成功!")
	}

	// mongoDb.AddOne(initMongo)
	// mongoDb.Count(initMongo)
	// mongoDb.GetList(bson.M{"level": 55}, initMongo)

	for _, file := range ScanDir() {
		fmt.Println(file)
		// paths, fileName := filepath.Split(file)
		res := ReadSettingsFromFile(file, &model.LogStruct{})
		fmt.Printf("%v", res)
		mongoDb.AddMany(initMongo, res)
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

func ReadSettingsFromFile(settingFilePath string, config *LogStruct) (arr []interface{}) {
	jsonFile, err := os.Open(settingFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	// byteValue, _ := ioutil.ReadAll(jsonFile)
	scanner := bufio.NewScanner(jsonFile)
	for scanner.Scan() {
		// var byteValue []byte
		byteValue := []byte(scanner.Text())
		fmt.Println("-------------")
		fmt.Printf("%s\n", byteValue)
		err = json.Unmarshal(byteValue, &config)
		arr = append(arr, *config)
		// if err != nil {
		// log.Panic(err)
		// }
	}
	return arr
}

// not used
func Readfile(name, dir string) (LogData *LogStruct) {
	config := viper.New()
	// config.SetConfigName(name) // name of config file (without extension)
	config.SetConfigName("test.yaml") // name of config file (without extension)

	config.SetConfigType("yaml")                                            // REQUIRED if the config file does not have the extension in the name
	config.AddConfigPath("/home/zxz/dnmp/www/rrzuji/console/logV3/202203/") // optionally look for config in the working directory
	err := config.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := config.Unmarshal(LogData); err != nil {
		fmt.Println(err)
	}
	return LogData
}

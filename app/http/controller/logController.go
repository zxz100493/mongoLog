package controller

import (
	"app-log/app/model"
	mongoDb "app-log/pkg/database/mongoDb"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
		// paths, fileName := filepath.Split(file)
		// extension := path.Ext(file)
		// // fileName = strings.Replace(fileName, extension, "", -1)
		// fmt.Println(fileName)
		// fmt.Println(extension)
		// fmt.Println(paths)

		// Readfile(fileName, paths)
		res := ReadSettingsFromFile(file)
		fmt.Printf("%v", res)
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

func ReadSettingsFromFile(settingFilePath string) (config *LogStruct) {
	jsonFile, err := os.Open(settingFilePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Printf("%s", byteValue)
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Panic(err)
	}
	return config
}

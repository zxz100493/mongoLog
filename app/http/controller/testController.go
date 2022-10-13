package controller

// import (
// 	"context"
// 	"log"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/mongo/readpref"
// )

// // 数据结构体
// type Test struct {
// 	Id    string `bson:"_id"`
// 	Name  string `bson:"name"`
// 	Level int    `bson:"level"`
// }

// var collection *mongo.Collection // collection 话柄

// func TestLog(c *gin.Context) {
// 	uri := "mongodb://root:123456@127.0.0.1:27017/admin"
// 	name := "Test"
// 	maxTime := time.Duration(2) // 链接超时时间
// 	table := "test"             // 表名

// 	db, err := ConnectToDB(uri, name, maxTime)
// 	if err != nil {
// 		panic("链接数据库有误!")
// 	}

// 	collection = db.Collection(table)

// 	t := Test{
// 		Id:    "1",
// 		Name:  "zngw",
// 		Level: 55,
// 	}
// 	// 添加一条数据
// 	AddOne(&t)
// 	Count()
// }

// func Count() {
// 	count, err := collection.CountDocuments(context.Background(), bson.D{})
// 	if err != nil {
// 		log.Fatal(count)
// 	}
// 	log.Println("collection.CountDocuments:", count)
// }

// func ConnectToDB(uri, name string, timeout time.Duration) (*mongo.Database, error) {
// 	// 设置连接超时时间
// 	ctx, cancel := context.WithTimeout(context.Background(), timeout)
// 	defer cancel()
// 	// 通过传进来的uri连接相关的配置
// 	o := options.Client().ApplyURI(uri)
// 	// 发起链接
// 	client, err := mongo.Connect(ctx, o)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	// 判断服务是不是可用
// 	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	// 返回 client
// 	return client.Database(name), nil
// }

// func AddOne(t *Test) {
// 	objId, err := collection.InsertOne(context.TODO(), &t)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	log.Println("录入数据成功，objId:", objId)
// }

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// type Post struct { //带结构标签，反引号来包围字符串
// 	Id      int       `json:"id"`
// 	Content string    `json:"content"`
// 	Author  Author    `json:"author"`
// 	Comment []Comment `json:"comments"`
// }

// type Author struct {
// 	Id   int    `json:"id"`
// 	Name string `json:"name"`
// }

// type Comment struct {
// 	Id      int    `json:"id"`
// 	Content string `json:"content"`
// 	Author  string `json:"author"`
// }

// func CD(c *gin.Context) {
// 	jsonFile, err := os.Open("//home/zxz/gogo/go-mongo-log/app/http/controller/test.json")
// 	if err != nil {
// 		fmt.Println("error opening json file")
// 		return
// 	}
// 	defer jsonFile.Close()

// 	jsonData, err := ioutil.ReadAll(jsonFile)
// 	fmt.Printf("%s", jsonData)
// 	if err != nil {
// 		fmt.Println("error reading json file")
// 		return
// 	}
// 	var post Post
// 	json.Unmarshal(jsonData, &post)
// 	fmt.Println(post)
// }

type Comment struct {
	Ccontent ContextContent `json:"context"`
}

type ContextContent struct {
	NotifyType        string `json:"notify_type"`
	PayStatus         string `json:"pay_status"`
	NotifyAppId       string `json:"notify_app_id"`
	AlipayFundOrderNo string `json:"alipay_fund_order_no"`
	PayAmount         string `json:"pay_amount"`
	ZmOrderNo         string `json:"zm_order_no"`
	OutTransNo        string `json:"out_trans_no"`
	Channels          string `json:"channel"`
	OutOrderNo        string `json:"out_order_no"`
	PayTime           string `json:"pay_time"`
	SignType          string `json:"sign_type"`
	Sign              string `json:"sign"`
}

func CD(c *gin.Context) {
	jsonFile, err := os.Open("/home/zxz/codego/mongoLog/app/http/controller/test.json")
	if err != nil {
		fmt.Println("error opening json file")
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	// fmt.Printf("%s", jsonData)
	if err != nil {
		fmt.Println("error reading json file")
		return
	}
	var comment []Comment
	json.Unmarshal(jsonData, &comment)
	// fmt.Println(comment)
	var num int
	for _, v := range comment {
		s := v
		// fmt.Println(s.Ccontent.SignType)
		// return
		data := make(url.Values)
		data["notify_type"] = []string{s.Ccontent.NotifyType}
		data["pay_status"] = []string{s.Ccontent.PayStatus}
		data["notify_app_id"] = []string{s.Ccontent.NotifyAppId}
		data["alipay_fund_order_no"] = []string{s.Ccontent.AlipayFundOrderNo}
		data["pay_amount"] = []string{s.Ccontent.PayAmount}
		data["zm_order_no"] = []string{s.Ccontent.ZmOrderNo}
		data["out_trans_no"] = []string{s.Ccontent.OutTransNo}
		data["channel"] = []string{s.Ccontent.Channels}
		data["out_order_no"] = []string{s.Ccontent.OutOrderNo}
		data["pay_time"] = []string{s.Ccontent.PayTime}
		data["sign"] = []string{s.Ccontent.Sign}
		data["sign_type"] = []string{"RSA2"}
		// if s.Ccontent.OutOrderNo != "2020032614474456" {
		// 	continue
		// }
		time.Sleep(3 * time.Second)
		// fmt.Println(data)
		// fmt.Println(s.Ccontent.SignType)
		// res, _ := http.PostForm("http://m.rrzuji.com/individual/api/notify", data)
		res, _ := http.PostForm("", data)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Printf("%s", body)
		num++
		fmt.Println(num)

		fmt.Println("")
	}
	fmt.Println("DONE")
}

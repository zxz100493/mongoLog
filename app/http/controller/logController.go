package controller

import (
	mongoDb "app-log/pkg/database/mongoDb"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func TestC(c *gin.Context) {
	uri := "mongodb://root:123456@127.0.0.1:27017/admin"
	name := "Test"
	maxTime := time.Duration(2) // 链接超时时间
	// table := "test"             // 表名

	_, err := mongoDb.ConnectToDB(uri, name, maxTime)

	if err != nil {
		log.Println("链接数据库有误!")
	} else {
		log.Println("链接成功!")
	}
	mongoDb.Count()
}

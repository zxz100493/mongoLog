package mongoDB

import (
	"app-log/config"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectToDB(name string) (*mongo.Database, error) {
	MongoConfig := config.Instance.MongoDB
	user := MongoConfig.User
	port := MongoConfig.Port
	host := MongoConfig.Host
	pwd := MongoConfig.Pwd
	timeout := time.Duration(2) // 链接超时时间

	// uri := fmt.Sprintf("mongodb://%s:123456@127.0.0.1:27017/admin", config.Instance.Mysql.User)
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/admin", user, pwd, host, port)
	// 设置连接超时时间
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	// 通过传进来的uri连接相关的配置
	o := options.Client().ApplyURI(uri)
	// 发起链接
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	// 判断服务是不是可用
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
		return nil, err
	}
	// 返回 client
	return client.Database(name), nil
}

package controller

import (
	mongoDB "app-log/pkg/database/mongoDb"
	tools "app-log/pkg/tools/json"
	"context"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

type mongoObject struct {
	Client *mongo.Client
}

type mongoMemory struct {
	bits     int `bson."bits"`
	resident int `bson."resident"`
	virtual  int `bson."virtual"`
}

type hostInfo struct {
	uptime   uint64 `json:"uptime"`
	platform uint64 `json:"platform"`
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

func GetConn() {
	if conn == nil {
		once.Do(func() {
			conn = NewMongoObject()
		})
	}
}

func GetSystemInfo(c *gin.Context) {
	GetConn()
	// dbs, err := conn.Client.Version()
	// dbs, err := conn.Client.ListDatabaseNames(c, bson.M{})
	// ConnectToDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("%v", dbs)
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	serverStatus, err := conn.Client.Database("admin").RunCommand(
		ctx,
		bsonx.Doc{{"serverStatus", bsonx.Int32(1)}},
	).DecodeBytes()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(serverStatus)
	fmt.Println(reflect.TypeOf(serverStatus))
	version, err := serverStatus.LookupErr("version")
	if err != nil {
		fmt.Println(err)
	}

	uptime, err := serverStatus.LookupErr("uptime")
	if err != nil {
		fmt.Println(err)
	}

	memory, err := serverStatus.LookupErr("mem")
	if err != nil {
		fmt.Println(err)
	}

	memoryInfo, err := memory.Document().LookupErr("bits")
	if err != nil {
		fmt.Println(err)
	}
	virtualMemoryInfo, err := memory.Document().LookupErr("virtual")
	if err != nil {
		fmt.Println(err)
	}
	data := make(map[string]interface{})
	data["version"] = version.StringValue()
	data["uptime"] = uptime.Double()
	data["memory"] = memoryInfo.Int32()
	data["virtualMemory"] = virtualMemoryInfo.Int32()

	gethostInfo(data)
	c.JSON(200, gin.H{"msg": "ok", "status": tools.SUCCESS, "data": data})
}

func gethostInfo(data map[string]interface{}) {
	host, _ := host.Info()
	cpuNum, _ := cpu.Counts(true)
	memInfo, _ := mem.VirtualMemory()
	data["cpuNum"] = cpuNum
	data["platform"] = host.Platform
	data["hostUptime"] = host.Uptime
	data["mem"] = formatFileSize(memInfo.Total)
	data["memUsed"] = formatFileSize(memInfo.Used)
}

func formatFileSize(fileSize uint64) (size string) {
	if fileSize < 1024 {
		//return strconv.FormatInt(fileSize, 10) + "B"
		return fmt.Sprintf("%.2fB", float64(fileSize)/float64(1))
	} else if fileSize < (1024 * 1024) {
		return fmt.Sprintf("%.2fKB", float64(fileSize)/float64(1024))
	} else if fileSize < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fMB", float64(fileSize)/float64(1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fGB", float64(fileSize)/float64(1024*1024*1024))
	} else if fileSize < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2fTB", float64(fileSize)/float64(1024*1024*1024*1024))
	} else { //if fileSize < (1024 * 1024 * 1024 * 1024 * 1024 * 1024)
		return fmt.Sprintf("%.2fEB", float64(fileSize)/float64(1024*1024*1024*1024*1024))
	}
}

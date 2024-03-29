package controller

import (
	"app-log/app/service"
	tools "app-log/pkg/tools/json"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

func GetSystemInfo(c *gin.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	serverStatus, err := service.Conn.Client.Database("admin").RunCommand(
		ctx,
		bsonx.Doc{{"serverStatus", bsonx.Int32(1)}},
	).DecodeBytes()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(serverStatus)
	// fmt.Println(reflect.TypeOf(serverStatus))
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

func GetDbNameList(c *gin.Context) {
	dbs, err := service.Conn.Client.ListDatabaseNames(c, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(200, gin.H{"msg": "ok", "status": tools.SUCCESS, "data": dbs})
}

func GetDbDetail(c *gin.Context) {
	name := c.Query("name")
	cls, err := service.Conn.Client.Database(name).ListCollectionNames(c, bson.M{})
	if err != nil {
		fmt.Println(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	var document bson.M
	err = service.Conn.Client.Database(name).RunCommand(
		ctx,
		bsonx.Doc{{"dbStats", bsonx.Int32(1)}},
	).Decode(&document)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(document)

	ret := make(map[string]interface{}, 2)
	ret["stat"] = document
	ret["cls"] = cls

	c.JSON(200, gin.H{"msg": "ok", "status": tools.SUCCESS, "data": ret})
}

func GetClsDetail(c *gin.Context) {
	name := c.Query("name")
	log.Printf("name:%s,end", name)
	var document bson.M
	err := service.Conn.Client.Database(name).RunCommand(
		context.Background(),
		bson.M{"collStats": name},
	).Decode(&document)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(document)

	ret := make(map[string]interface{}, 2)
	ret["stat"] = document

	c.JSON(200, gin.H{"msg": "ok", "status": tools.SUCCESS, "data": ret})
}

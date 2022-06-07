package main

import (
	"app-log/config"
	"app-log/router"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r := gin.New()
	//r.Use(middleware.StatCount())

	router.LoadLog(r)
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config.Init(fmt.Sprintf("%s/../", path))

	if err := r.Run(":8888"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}

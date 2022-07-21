package main

import (
	"app-log/config"
	"app-log/router"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()
	r := gin.New()
	r.Use(cors.Default())

	router.LoadLog(r)
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	config.Init(fmt.Sprintf("%s/", path))
	// fmt.Printf("%s/../", path)
	// fmt.Printf("%s/../", path)
	if err := r.Run(":8888"); err != nil {
		fmt.Printf("startup service failed, err:%s\n", err)
	}
	fmt.Println("startup service")
}

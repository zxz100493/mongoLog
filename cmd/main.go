package main

import (
	"app-log/app/service"
	"app-log/config"
	"app-log/router"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default()
	r := gin.New()
	r.Use(cors.Default())
	// r.StaticFile("/", "./resource/dist/index.html")

	router.LoadLog(r)
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	r.GET("/", gin.WrapH(http.FileServer(http.Dir("./resource/dist")))) // chargpt recommended

	config.Init(fmt.Sprintf("%s/", path))
	service.GetConn()
	// fmt.Printf("%s/../", path)
	// fmt.Printf("%s/../", path)
	if err := r.Run(":8888"); err != nil {
		fmt.Printf("startup service failed, err:%s\n", err)
	}
	fmt.Println("startup service")
}

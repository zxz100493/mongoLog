package main

import (
	"app-log/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r := gin.New()
	//r.Use(middleware.StatCount())
	router.LoadLog(r)

	if err := r.Run(":8888"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}

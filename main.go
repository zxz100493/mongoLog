package main

import (
	"app-log/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.创建路由
	// r := gin.Default()
	// // 2.绑定路由规则，执行的函数
	// routers.LoadTest(r)
	// // 3.监听端口，默认在8080
	// // Run("里面不指定端口号默认为8080")
	// r.Run(":8888")
	r := gin.Default()
	//r := gin.New()
	//r.Use(middleware.StatCount())
	router.LoadTest(r)
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}

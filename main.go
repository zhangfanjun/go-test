package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/gops/agent"
	_ "net/http/pprof"
	"test/src/controller"
)

func init() {
	fmt.Print("初始化main包的数据。。。。。")

}

func main() {
	if err := agent.Listen(agent.Options{Addr: ":8899"}); err != nil {
		fmt.Print("启动监听失败", err)
	}
	engine := gin.Default()
	//开放静态资源的访问
	engine.Static("/static", "/resource/static")
	//接口路由
	controller.TestController(engine)
	//服务器的端口号
	var port uint64 = 8080
	serverPort := fmt.Sprintf(":%v", port)
	err := engine.Run(serverPort)
	if err != nil {
		fmt.Print("启动失败")
	}
}

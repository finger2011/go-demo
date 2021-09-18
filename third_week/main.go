package main

import (
	"fmt"
	// "net/http"
	"finger2011/third-week/web"
)

func home(c *web.Context) {
	fmt.Fprintf(c.W, "这是主页")
}

func user(c *web.Context) {
	fmt.Fprintf(c.W, "这是用户")
}

func createUser(c *web.Context) {
	fmt.Fprintf(c.W, "这是创建用户")
}

func order(c *web.Context) {
	fmt.Println("order called")
	fmt.Fprintf(c.W, "这是订单")
}

func main() {
	server := web.CreateSdkHTTPServer("test")
	fmt.Println("create server success!")
	//注册路由
	server.Route("GET", "/", home)
	server.Route("GET", "/user", user)
	server.Route("PUT", "/user", createUser)
	server.Route("GET", "/order", order)

	go func() {
		fmt.Println("start listen")
		if err := server.Start(":8080"); err != nil {
			// 启动失败，直接panic
			fmt.Printf("start error:%v", err)
			panic(err)
		}
		fmt.Println("start do something")
	}()

	// 关闭server
	// web.ShutDownPrepare()
	web.ShutDown()

	// select {}
}

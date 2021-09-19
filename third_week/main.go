package main

import (
	"context"
	"os"
	"syscall"
	"time"

	// "errors"
	"fmt"

	"golang.org/x/sync/errgroup"

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
	fmt.Fprintf(c.W, "这是订单")
}

//模拟系统调用退出
func mockSyscal(c *web.Context) {
	fmt.Fprintf(c.W, "模拟系统调用退出")
	ch <- syscall.SIGUSR1
}

//用于模拟系统调用，测试
var ch = make(chan os.Signal, 1)

//进程异常退出时，app退出
//TODO:也可以考虑放入App中
var closed = make(chan bool, 1)

func main() {
	app := web.CreateApp(context.Background(), "golang")
	group, ctx := errgroup.WithContext(app.GetContext())
	//多个server
	servers := []string{":8080", ":8081"}

	//启动多个server
	for index, serverName := range servers {
		serverName := serverName
		index := index
		var server = web.CreateSdkHTTPServer(ctx, serverName, closed)
		//路由
		if index == 0 {
			server.Route("GET", "/", home)
			server.Route("GET", "/user", user)
			server.Route("GET", "/syscall", mockSyscal)
		} else {
			server.Route("PUT", "/user", createUser)
			server.Route("GET", "/order", order)
		}
		group.Go(func() error {
			fmt.Println("start server" + serverName)
			if err := server.Start(ctx, serverName); err != nil {
				//启动失败，直接退出
				//Start永远返回nil，实际执行不到
				// return err
			}
			fmt.Println("start server done" + serverName)
			return nil
		})
		//把server挂载到App上，
		//TODO 这里可以考虑在创建app时，当做config传入
		app.AddServer(server)
	}
	//signal 信号的注册和处理
	// app.SetSignal([]os.Signal{}) 注册信号
	group.Go(func() error {
		return web.ShutDown(ctx, app, ch)
	})

	//测试：随机server shutdown掉
	// go testRandomShutDownServer(ctx, app)

	//server异常退出，app退出
	serverClosed(ctx, app)

	if err := group.Wait(); err != nil {
		fmt.Printf("server error:%v\n", err)
		app.Stop()
	}
}

//server异常退出时，app退出
func serverClosed(ctx context.Context, app *web.App) {
	go func() {
		select {
		case v, ok := <-closed:
			if ok && v {
				break
			}
		case <-ctx.Done():
			return
		}
		// fmt.Println("check server closed")
		app.Stop()
		return
	}()
}

//测试：随机server shutdown掉
func testRandomShutDownServer(ctx context.Context, app *web.App) {
	time.Sleep(time.Second * 5)
	app.RandomShutDownServer(ctx)
	return
}

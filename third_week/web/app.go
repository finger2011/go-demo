package web

import (
	"context"
	"fmt"
	"os"
	"time"
)

//App app config
type App struct {
	name    string
	ctx     context.Context
	cancel  func()
	servers []Server
}

//GetContext get app context
func (a App) GetContext() context.Context {
	return a.ctx
}

//GetName get app name
func (a App) GetName() string {
	return a.name
}

//AddServer add server
func (a *App) AddServer(s Server) {
	a.servers = append(a.servers, s)
}

//ShutDownServer shutdown server
func (a *App) ShutDownServer(ctx context.Context) {
	for _, server := range a.servers {
		server := server
		fmt.Println("shut down server")
		server.Stop(ctx)
	}
}

//RandomShutDownServer 随机关闭一个server，用于测试
func (a *App) RandomShutDownServer(ctx context.Context) {
	for _, server := range a.servers {
		server := server
		fmt.Println("random shut down server")
		server.Stop(ctx)
		break
	}
}

//Stop stop app
func (a App) Stop() error {
	fmt.Println("stop app...")
	//超时控制，超过1min未成功关闭，强制退出s
	time.AfterFunc(time.Minute, func() {
		fmt.Println("shut down timeout...")
		os.Exit(-1)
	})
	err := GracefllExit(a.ctx)
	if err != nil {
		//可能需要重试之类的机制
	}
	a.ShutDownServer(a.ctx)
	if a.cancel != nil {
		fmt.Println("cancel app context...")
		a.cancel()
		// fmt.Println("cancel app context done...")
	}
	// os.Exit(0)
	return nil
}

//CreateApp create app
func CreateApp(ctx context.Context, name string) *App {
	c, cancel := context.WithCancel(ctx)
	return &App{
		name:   name,
		ctx:    c,
		cancel: cancel,
	}
}

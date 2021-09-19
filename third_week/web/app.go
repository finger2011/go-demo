package web

import (
	"context"
	"fmt"
	"os"
	"syscall"
	"time"
)

//App app config
type App struct {
	name    string
	ctx     context.Context
	cancel  func()
	servers []Server
	sigs    []os.Signal
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

//SetSignal set signals
func (a *App) SetSignal(sigs []os.Signal) {
	a.sigs = sigs
}

//GetSignal get signals
func (a App) GetSignal() []os.Signal {
	return a.sigs
}

//ShutDownServer shutdown server
func (a *App) ShutDownServer(ctx context.Context) {
	for _, server := range a.servers {
		server := server
		fmt.Println("shut down server")
		server.Stop(ctx)
	}
	return
}

//RandomShutDownServer 随机关闭一个server，用于测试
func (a *App) RandomShutDownServer(ctx context.Context) {
	for _, server := range a.servers {
		server := server
		fmt.Println("random shut down server")
		server.Stop(ctx)
		break
	}
	return
}

//Stop stop app
func (a App) Stop() error {
	fmt.Println("stop app...")
	//超时控制，超过1min未成功关闭，强制退出
	time.AfterFunc(time.Minute, func() {
		fmt.Println("stop app timeout...")
		//TODO:是否有更好的方式?
		os.Exit(1)
	})
	err := GracefulExit(a.ctx)
	if err != nil {
		//TODO:可能需要重试之类的机制?
	}
	a.ShutDownServer(a.ctx)
	if a.cancel != nil {
		a.cancel()
	}
	fmt.Println("stop app done...")
	return nil
}

//CreateApp create app
func CreateApp(ctx context.Context, name string) *App {
	c, cancel := context.WithCancel(ctx)
	sigs := []os.Signal{syscall.SIGKILL, syscall.SIGINT, syscall.SIGUSR1, syscall.SIGBUS}
	return &App{
		name:   name,
		ctx:    c,
		cancel: cancel,
		sigs:   sigs,
	}
}

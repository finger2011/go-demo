package web

import (
	"context"
	"fmt"
	"net/http"
	// "time"
	// "golang.org/x/sync/errgroup"
)

//Routable 路由
type Routable interface {
	Route(method string, path string, handlerFunc handlerFunc)
}

//Server server
type Server interface {
	Routable

	Start(ctx context.Context, address string) error

	Stop(ctx context.Context) error
}

type sdkHTTPServer struct {
	srv       *http.Server
	Name      string
	handler   Handler
	closed    chan bool
	err       error
	srvClosed bool
}

func (s *sdkHTTPServer) Start(ctx context.Context, address string) error {
	s.srv = &http.Server{Addr: address, Handler: s.handler}
	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			s.err = err
			s.srvClosed = true
			s.closed <- true
			return
		}
	}()
	return nil
}

func (s *sdkHTTPServer) Route(method, pattern string, handlerFunc handlerFunc) {
	// http.HandleFunc(pattern, handlerFunc)
	s.handler.Route(method, pattern, handlerFunc)
}

func (s *sdkHTTPServer) Stop(ctx context.Context) error {
	fmt.Println("stop server" + s.Name)
	if s.srvClosed {
		fmt.Println("server already be closed" + s.Name)
		return nil
	}
	var err = s.srv.Shutdown(ctx)
	if err == nil {
		s.srvClosed = true
	}
	fmt.Println("stop server end" + s.Name)
	return err
}

//CreateSdkHTTPServer create a new server
func CreateSdkHTTPServer(ctx context.Context, name string, closed chan bool) Server {
	return &sdkHTTPServer{
		Name: name,
		handler: &HandlerRoute{
			handlers: make(map[string]func(con *Context)),
		},
		closed: closed,
	}
}

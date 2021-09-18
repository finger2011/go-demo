package web

import (
	"context"
	"net/http"
)

//Routable 路由
type Routable interface {
	Route(method string, path string, handlerFunc handlerFunc)
}

//Server server
type Server interface {
	Routable

	Start(address string) error

	Stop(ctx context.Context) error
}

type sdkHTTPServer struct {
	Name    string
	handler Handler
	root    Filter
}

func (s *sdkHTTPServer) Start(address string) error {
	// http.HandleFunc("/", func(writer http.ResponseWriter,
	// 	request *http.Request) {
	// 	c := CreateContext(writer, request)
	// 	s.root(c)
	// })
	return http.ListenAndServe(address, s.handler)
}

func (s *sdkHTTPServer) Route(method, pattern string, handlerFunc handlerFunc) {
	s.handler.Route(method, pattern, handlerFunc)
}

func (s *sdkHTTPServer) Stop(ctx context.Context) error {
	return nil
}

// func (s *sdkHTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	// c := s.ctxPool.Get().(*Context)
// 	// defer func() {
// 	// 	s.ctxPool.Put(c)
// 	// }()
// 	// c.Reset(writer, request)
// 	// s.root(ctx)
// }

//CreateSdkHTTPServer create a new server
func CreateSdkHTTPServer(name string) Server {
	return &sdkHTTPServer{
		Name: name,
		handler: &HandlerRoute{
			handlers: make(map[string]func(con *Context)),
		},
	}
}

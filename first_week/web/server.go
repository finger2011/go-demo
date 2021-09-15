package web

import (
	"net/http"
)

//Routable routable
type Routable interface {
	//Route bind route rules
	Route(method, pattern string, handlerFunc func(con *Context))
}

//Server interface
type Server interface {
	//Start server
	Start(address string) error

	Routable
}

type sdkHTTPServer struct {
	Name    string
	handler Handler
}

func (s *sdkHTTPServer) Start(address string) error {
	return http.ListenAndServe(address, s.handler)
}

func (s *sdkHTTPServer) Route(method, pattern string, handlerFunc func(con *Context)) {
	s.handler.Route(method, pattern, handlerFunc)
	// key := s.handler.key(method, pattern)
	// s.handler.handlers[key] = handlerFunc
}

//CreateSdkHTTPServer create a new server
func CreateSdkHTTPServer(name string) Server {
	return &sdkHTTPServer{
		Name: name,
		handler: &HandlerMap{
			handlers: make(map[string]func(con *Context)),
		},
	}
}

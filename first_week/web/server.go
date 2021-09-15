package web

import (
	"fmt"
	"net/http"
)

type commonResponse struct {
	BizCode int
	Msg string
}


type signupReq struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	Passward          string `json:"passward"`
	ConfirmedPassward string `json:"confirmed_passward"`
}

//Signup signup user
func Signup(con *Context) {
	var req = &signupReq{}
	err := con.ReadJSON(req)
	if err != nil {
		var resp = &commonResponse{
			BizCode: 1,
			Msg: fmt.Sprintf("invalid request:%v", err),
		}
		err = con.WriteJSON(1, resp)
		if err != nil {

		}
		return
	}
	fmt.Fprintf(con.W, "%d", err)
}

//Server interface
type Server interface {
	//Start server
	Start(address string) error

	//Route bind route rules
	Route(patter string, handlerFunc func(con *Context))
}

type sdkHTTPServer struct {
	Name string
}

func (s *sdkHTTPServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func (s *sdkHTTPServer) Route(patter string, handlerFunc func(con *Context)) {
	http.HandleFunc(patter, func (w http.ResponseWriter, r *http.Request)  {
		var con = CreateContext(w, r)
		handlerFunc(con)
	})
}

//CreateSdkHTTPServer create a new server
func CreateSdkHTTPServer(name string) Server {
	return &sdkHTTPServer{
		Name: name,
	}
}

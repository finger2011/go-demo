package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type commonResponse struct {
	BizCode int
	Msg string
}

//Context w r
type Context struct {
	W http.ResponseWriter
	R *http.Request
}

//ReadJSON read json
func (con *Context) ReadJSON(data interface{}) error {
	body, err := io.ReadAll(con.R.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}

type signupReq struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	Passward          string `json:"passward"`
	ConfirmedPassward string `json:"confirmed_passward"`
}

//Signup signup user
func Signup(w http.ResponseWriter, r *http.Request) {
	var req = &signupReq{}
	var con = Context{
		W: w,
		R: r,
	}
	err := con.ReadJSON(req)
	if err != nil {
		var resp = &commonResponse{
			BizCode: 1,
			Msg: fmt.Sprintf("invalid request:%v", err),
		}
		respBytes, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(respBytes))
		return
	}
	fmt.Fprintf(w, "%d", err)
}

//Server interface
type Server interface {
	//Start server
	Start(address string) error

	//Route bind route rules
	Route(patter string, handlerFunc http.HandlerFunc)
}

type sdkHTTPServer struct {
	Name string
}

func (s *sdkHTTPServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func (s *sdkHTTPServer) Route(patter string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(patter, handlerFunc)
}

//CreateSdkHTTPServer create a new server
func CreateSdkHTTPServer(name string) Server {
	return &sdkHTTPServer{
		Name: name,
	}
}

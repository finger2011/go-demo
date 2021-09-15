package main

import (
	"finger2011/first-week/web"
	"fmt"
)

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type signupReq struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	Passward          string `json:"passward"`
	ConfirmedPassward string `json:"confirmed_passward"`
}

//Signup signup user
func Signup(con *web.Context) {
	var req = &signupReq{}
	err := con.ReadJSON(req)
	if err != nil {
		var resp = &commonResponse{
			BizCode: 1,
			Msg:     fmt.Sprintf("invalid request:%v", err),
		}
		err = con.WriteJSON(1, resp)
		if err != nil {
			fmt.Fprintf(con.W, "write json error:%v", err)
		}
		return
	}
	// fmt.Fprintf(con.W, "%d", err)
	con.WriteJSON(200, &commonResponse{
		Data: 123,
	})
}

package main

import (
	"finger2011/first-week/web"
	"fmt"
)

func handler(con *web.Context) {
	fmt.Fprintf(con.W, "Hello, %s", con.R.URL.Path[1:])
}

func home(con *web.Context) {
	fmt.Fprintf(con.W, "This is home path: %s\n", con.R.URL.Path[1:])
	queryParams(con)
	getHTTPHost(con)
	getHTTPHeader(con)
	getHTTPForm(con)
}

func user(con *web.Context) {
	fmt.Fprintf(con.W, "This is user path\n")
	readHTTPBody(con)
}

// func createUser(con.W http.ResponseWriter, con.R *http.Request) {
// 	fmt.Fprintf(con.W, "This is create user path\n")
// 	var con = web.Context{
// 		W: con.W,
// 		R: con.R,
// 	}
// 	web.Signup(con)
// }

func order(con *web.Context) {
	fmt.Fprintf(con.W, "This is order path\n")
	getHTTPBody(con)
}

func main() {
	server := web.CreateSdkHTTPServer("test-server")
	server.Route("GET","/", home)
	server.Route("GET","/user", user)
	server.Route("PUT","/user/create", Signup)
	server.Route("GET","/order", order)
	server.Start(":8080")
}

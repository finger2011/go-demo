package main

import (
	"finger2011/first-week/web"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is home path: %s\n", r.URL.Path[1:])
	queryParams(w, r)
	getHTTPHost(w, r)
	getHTTPHeader(w, r)
	getHTTPForm(w, r)
}

func user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is user path\n")
	readHTTPBody(w, r)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is create user path\n")
	web.Signup(w, r)
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is order path\n")
	getHTTPBody(w, r)
}

func main() {
	server := web.CreateSdkHTTPServer("test-server")
	server.Route("/", home)
	server.Route("/user", user)
	server.Route("/user/create", createUser)
	server.Route("/order", order)
	server.Start(":8080")
}

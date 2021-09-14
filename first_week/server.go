package main

import (
	"fmt"
	"log"
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
	fmt.Fprintf(w, "This is user path: %s\n", r.URL.Path[2:])
	readHTTPBody(w, r)
}

func order(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is order path: %s\n", r.URL.Path[2:])
	getHTTPBody(w, r)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/user", user)
	http.HandleFunc("/order", order)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

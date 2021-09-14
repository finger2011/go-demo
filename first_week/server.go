package main

import (
	"log"
	"net/http"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello, %s", r.URL.Path[1:])
}


func home(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "This is home path: %s", r.URL.Path[1:])
}

func user(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "This is user path: %s", r.URL.Path[2:])
}

func order(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "This is order path: %s", r.URL.Path[2:])
}

func main()  {
	http.HandleFunc("/", home)
	http.HandleFunc("/user", user)
	http.HandleFunc("/order", order)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
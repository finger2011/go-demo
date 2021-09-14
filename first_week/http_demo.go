package main

import (
	"fmt"
	"io"
	"net/http"
)

//ReadHTTPBody test function
func readHTTPBody(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "first read body failed: %v\n", err)
		return
	}
	fmt.Fprintf(w, "first data from body: %v\n", body)

	body, err = io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "second read body failed: %v\n", err)
		return
	}
	fmt.Fprintf(w, "second data from body: %s; and data length:%d\n", string(body), len(body))
}

func getHTTPBody(w http.ResponseWriter, r *http.Request) {
	if r.GetBody == nil {
		fmt.Fprintf(w, "get body is nil\n")
	} else {
		fmt.Fprintf(w, "get body is not nil\n")
	}
}

func queryParams(w http.ResponseWriter, r *http.Request) {
	uri := r.URL.Query()
	fmt.Fprintf(w, "query param:%v\n", uri)
}

func getHTTPHost(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	fmt.Fprintf(w, "host:%s\n", host)
}

func getHTTPHeader(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	fmt.Fprintf(w, "host:%s\n", header)
}

func getHTTPForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "before parse form:%v\n", r.Form)
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "parse form error:%v\n", err)
	}
	fmt.Fprintf(w, "after parse form:%v\n", r.Form)
}

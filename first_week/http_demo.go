package main

import (
	"finger2011/first-week/web"
	"fmt"
	"io"
)

//ReadHTTPBody test function
func readHTTPBody(con *web.Context) {
	body, err := io.ReadAll(con.R.Body)
	if err != nil {
		fmt.Fprintf(con.W, "first read body failed: %v\n", err)
		return
	}
	fmt.Fprintf(con.W, "first data from body: %v\n", body)

	body, err = io.ReadAll(con.R.Body)

	if err != nil {
		fmt.Fprintf(con.W, "second read body failed: %v\n", err)
		return
	}
	fmt.Fprintf(con.W, "second data from body: %s; and data length:%d\n", string(body), len(body))
}

func getHTTPBody(con *web.Context) {
	if con.R.GetBody == nil {
		fmt.Fprintf(con.W, "get body is nil\n")
	} else {
		fmt.Fprintf(con.W, "get body is not nil\n")
	}
}

func queryParams(con *web.Context) {
	uri := con.R.URL.Query()
	fmt.Fprintf(con.W, "query param:%v\n", uri)
}

func getHTTPHost(con *web.Context) {
	host := con.R.Host
	fmt.Fprintf(con.W, "host:%s\n", host)
}

func getHTTPHeader(con *web.Context) {
	header := con.R.Header
	fmt.Fprintf(con.W, "host:%s\n", header)
}

func getHTTPForm(con *web.Context) {
	fmt.Fprintf(con.W, "before parse form:%v\n", con.R.Form)
	err := con.R.ParseForm()
	if err != nil {
		fmt.Fprintf(con.W, "parse form error:%v\n", err)
	}
	fmt.Fprintf(con.W, "after parse form:%v\n", con.R.Form)
}

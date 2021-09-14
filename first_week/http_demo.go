package main

import (
	"fmt"
	"io"
	"net/http"
)

var a ="aaa"

//ReadHTTPBody test function
func readHTTPBody(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "\n first read body failed: %v", err)
		return
	}
	fmt.Fprintf(w, "\n first data from body: %v", body)

	body, err = io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "\n second read body failed: %v", err)
		return
	}
	fmt.Fprintf(w, "\n second data from body: %s; and data length:%d", string(body), len(body))
}

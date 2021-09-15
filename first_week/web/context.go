package web

import (
	"encoding/json"
	"io"
	"net/http"
)

//CreateContext create new context 
func CreateContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W: w,
		R: r,
	}
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

//WriteJSON write json
func (con *Context) WriteJSON(code int, data interface{}) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = con.W.Write(bs)
	if err != nil {
		return err
	}
	con.W.WriteHeader(code)
	return nil
}
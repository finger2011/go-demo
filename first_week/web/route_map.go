package web

import (
	"fmt"
	"net/http"
)

var _ Handler = &HandlerMap{}

//HandlerMap route map
type HandlerMap struct {
	handlers map[string]func(con *Context)
}

func (h *HandlerMap) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := h.key(r.Method, r.URL.Path)
	if handler, ok := h.handlers[key]; ok {
		con := CreateContext(w, r)
		handler(con)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("can not found any match route"))
	}
}

//Route routable
func (h *HandlerMap) Route(method string, pattern string, handlerFunc func(con *Context))  {
	key := h.key(method, pattern)
	h.handlers[key] = handlerFunc
}

func (h *HandlerMap) key(method, path string) string {
	return fmt.Sprintf("%s##%s", method, path)
}

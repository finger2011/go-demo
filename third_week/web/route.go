package web

import (
	"fmt"
	"net/http"
)

var _ Handler = &HandlerRoute{}

//HandlerRoute 路由决策
type HandlerRoute struct {
	handlers map[string]func(con *Context)
}

func (h *HandlerRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
func (h *HandlerRoute) Route(method string, pattern string, handlerFunc handlerFunc) {
	key := h.key(method, pattern)
	h.handlers[key] = handlerFunc
}

func (h *HandlerRoute) key(method, path string) string {
	return fmt.Sprintf("%s##%s", method, path)
}

package web

import (
	"fmt"
)

//HandlerTree binary tree
type HandlerTree struct {
	name     string
	Children []HandlerTree
}

func (h *HandlerTree) ServeHTTP(con *Context) {

}

//Route routable
func (h *HandlerTree) Route(method string, pattern string, handlerFunc func(con *Context)) {
	// key := h.key(method, pattern)
	// h.handlers[key] = handlerFunc
}

func (h *HandlerTree) key(method, path string) string {
	return fmt.Sprintf("%s##%s", method, path)
}

package web

import (
	"net/http"
)

//Handler handler
type Handler interface {
	http.Handler
	Routable
}

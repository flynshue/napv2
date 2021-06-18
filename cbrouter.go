package napv2

import (
	"fmt"
	"net/http"
)

type RouterFunc func(resp *http.Response) error

type CBRouter struct {
	Routers       map[int]RouterFunc
	DefaultRouter RouterFunc
}

func NewRouter() *CBRouter {
	return &CBRouter{
		Routers: make(map[int]RouterFunc),
		DefaultRouter: func(resp *http.Response) error {
			return fmt.Errorf("%d unknown status from %s", resp.StatusCode, resp.Request.URL.Path)
		},
	}
}

func (r *CBRouter) RegisterFunc(status int, handler RouterFunc) {
	r.Routers[status] = handler
}

func (r *CBRouter) CallFunc(resp *http.Response) error {
	f, ok := r.Routers[resp.StatusCode]
	if !ok {
		f = r.DefaultRouter
	}
	return f(resp)
}

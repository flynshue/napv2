package napv2

import "net/http"

type RouterFunc func(resp *http.Response) error

type CBRouter struct {
	Routers       map[int]RouterFunc
	DefaultRouter RouterFunc
}

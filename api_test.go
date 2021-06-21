package napv2

import (
	"net/http"
	"testing"
)

func TestAPI_Call(t *testing.T) {
	api := NewAPI("https://httpbin.org")
	router := NewRouter()
	router.RegisterFunc(200, func(resp *http.Response) error {
		return nil
	})
	resource := NewResource("GET", "/get", router)
	api.AddResource("get", resource)
	if err := api.Call("get", nil, nil); err != nil {
		t.Fail()
	}
}

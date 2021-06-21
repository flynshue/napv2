package napv2

import (
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {
	client := NewClient()
	baseURL := "https://httpbin.org"
	router := NewRouter()
	router.RegisterFunc(200, func(resp *http.Response) error {
		return nil
	})
	resource := NewResource("GET", "/get", router)
	if err := client.ProcessRequest(baseURL, resource, nil, nil); err != nil {
		t.Fail()
	}
}

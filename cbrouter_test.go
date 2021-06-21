package napv2

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestCBRouter(t *testing.T) {
	router := NewRouter()
	router.RegisterFunc(200, func(resp *http.Response) error {
		return nil
	})
	testURL, err := url.Parse("https://httpbin.org/get")
	if err != nil {
		t.FailNow()
	}
	resp := &http.Response{
		StatusCode: 200,
		Request: &http.Request{
			URL: testURL,
		},
	}
	if err := router.CallFunc(resp); err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

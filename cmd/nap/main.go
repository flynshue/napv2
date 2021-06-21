package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/flynshuePersonal/napv2"
)

var api = napv2.NewAPI("https://httpbin.org")

func main() {
	list := flag.Bool("list", false, "list available api resources")
	flag.Parse()
	if *list {
		for _, name := range api.ListResources() {
			fmt.Println(name)
			return
		}
	}
	if err := api.Call("get", nil, nil); err != nil {
		log.Fatalln("error api.Call(): ", err)
	}
}

func init() {
	router := napv2.NewRouter()
	router.RegisterFunc(200, func(resp *http.Response) error {
		defer resp.Body.Close()
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	})
	resource := napv2.NewResource("GET", "/get", router)
	api.AddResource("get", resource)
}

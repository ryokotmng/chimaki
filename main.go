package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	endpoint   = flag.String("endpoint", "", "endpoint url this tool will load")
	httpMethod = flag.String("http_method", "", "http request method")
)

func main() {
	flag.Parse()
	sendRequest(*endpoint)
}

func sendRequest(endPoint string) {
	switch *httpMethod {
	case "PUT":
		fmt.Println("unimplemented!")
	default:
		http.Get(endPoint)
	}
}

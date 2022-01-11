package main

import (
	"flag"
	"net/http"
)

var (
	endpoint = flag.String("endpoint", "", "endpoint url this tool will load")
)

func main() {
	flag.Parse()
	sendRequest(*endpoint)
}

func sendRequest(endPoint string) {
	http.Get(endPoint)
}

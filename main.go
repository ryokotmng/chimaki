package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var (
	endpoint   = flag.String("endpoint", "", "endpoint url this tool will load")
	httpMethod = flag.String("http_method", "", "http request method")
	interval   = flag.Uint64("interval", 1000, "interval(milliseconds) to load. Default value is 1,000 and must be uint64")
	duration   = flag.Uint64("duration", 60, "duration(seconds) to load. Default value is 60 and must be uint64")
)

func main() {
	flag.Parse()
	cancel := make(chan struct{})
	go sendRequest(cancel, *endpoint)
	time.Sleep(time.Duration(*duration) * time.Second)
	close(cancel)
}

func sendRequest(cancel chan struct{}, endPoint string) {
	t := time.NewTicker(time.Duration(*interval) * time.Millisecond)
	for {
		select {
		case <-cancel:
			return
		case <-t.C:
			switch *httpMethod {
			case "PUT":
				fmt.Println("unimplemented!")
			default:
				http.Get(endPoint)
			}
		}
	}
	t.Stop()
}

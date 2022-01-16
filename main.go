package main

import (
	"flag"
	"fmt"
	"log"
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

func sendRequest(cancel chan struct{}, endpoint string) {
	t := time.NewTicker(time.Duration(*interval) * time.Millisecond)
	testStartTime := time.Now()
	var numOfRequestsSent int
	for {
		select {
		case <-cancel:
			return
		case <-t.C:
			requestStartTime := time.Now()
			switch *httpMethod {
			case "PUT":
				fmt.Println("unimplemented!")
			default:
				res, err := http.Get(endpoint)
				if err != nil {
					break
				}
				if err := res.Body.Close(); err != nil {
					log.Fatalf("failed to close response body: %+v", err)
				}
				printDetails(numOfRequestsSent, time.Since(testStartTime).Milliseconds(), time.Since(requestStartTime).Milliseconds(), res)
			}
		}
		numOfRequestsSent++
	}
	t.Stop()
}

func printDetails(numOfRequestsSent int, elapsedTime int64, latency int64, res *http.Response) {
	fmt.Printf("number of requests sent: %v\n", numOfRequestsSent)
	fmt.Printf("elapsed time of the test: %vms\n", elapsedTime)
	fmt.Printf("latency: %vms\n", latency)
	fmt.Printf("response(status code: %v) ", res.StatusCode)
}

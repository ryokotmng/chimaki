package main

import (
	"context"
	"flag"
	"time"
)

var (
	endpoint   = flag.String("endpoint", "", "endpoint url this tool will load")
	httpMethod = flag.String("http_method", "", "http request method")
	duration   = flag.Uint64("duration", 60, "duration(seconds) to load. Default value is 60 and must be uint64")
	rate       = flag.Uint64("rate", 50, "Number of requests per time unit [0 = infinity] (default 50/1s)")
)

func main() {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := newClient(*endpoint)
	go c.sendRequest(ctx)
	time.Sleep(time.Duration(*duration) * time.Second)
}

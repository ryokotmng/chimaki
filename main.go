package main

import (
	"context"
	"flag"
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := newClient(*endpoint)
	go c.sendRequest(ctx)
	time.Sleep(time.Duration(*duration) * time.Second)
}

package main

import (
	"chimaki/lib"
	"context"
	"flag"
	"time"
)

var (
	url        = flag.String("url", "http://localhost:8080/", "url url this tool will load")
	httpMethod = flag.String("http_method", "GET", "http request method")
	duration   = flag.Uint64("duration", 60, "duration(seconds) to load. Default value is 60 and must be uint64")
	rate       = flag.Uint64("rate", 50, "Number of requests per time unit [0 = infinity] (default 50/1s)")
)

func main() {
	flag.Parse()
	ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	c := chimaki.NewClient(*url, *httpMethod, *rate)
	go c.SendRequest(ctx)
	time.Sleep(time.Duration(*duration) * time.Second)
	cancel()
	// TODO: use channel to stop the main thread
	time.Sleep(time.Duration(*duration) * 10 * time.Millisecond)
}

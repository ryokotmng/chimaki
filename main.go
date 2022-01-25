package main

import (
	"chimaki/lib"
	"context"
	"flag"
	"time"
)

var (
	duration = flag.Uint64("duration", 60, "duration(seconds) to load. Default value is 60 and must be uint64")
)

var options chimaki.ClientOptions

func init() {
	flag.StringVar(&options.Url, "url", "http://localhost:8080/", "url url this tool will load")
	flag.StringVar(&options.HttpMethod, "http_method", "GET", "http request method")
	flag.Uint64Var(&options.Rate, "rate", 50, "number of requests per time unit [0 = infinity] (default 50/1s)")
	flag.StringVar(&options.Bodyf, "body_file", "", "the file path which contains text to body")
}

func main() {
	flag.Parse()
	timeout := time.Duration(*duration) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	c := chimaki.NewClient(&options)
	results := c.ExecuteLoadTest(ctx)
	results.CreateMetrics()
}

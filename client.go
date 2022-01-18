package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type client struct {
	endpoint string
	*http.Client
}

func newClient(endpoint string) *client {
	c := *http.DefaultClient
	clone := &c
	return &client{endpoint, clone}
}

func (c *client) sendRequest(ctx context.Context) {
	req, err := http.NewRequest(*httpMethod, c.endpoint, nil)
	if err != nil {
		return
	}
	t := time.NewTicker(time.Duration(1000 / *rate) * time.Millisecond)
	testStartTime := time.Now()
	var numOfRequestsSent int
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			requestStartTime := time.Now()
			res, err := c.Do(req)
			if err != nil {
				break
			}
			if err := res.Body.Close(); err != nil {
				log.Fatalf("failed to close response body: %+v", err)
			}
			printDetails(numOfRequestsSent, time.Since(testStartTime).Milliseconds(), time.Since(requestStartTime).Milliseconds(), res)
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

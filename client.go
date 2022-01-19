package main

import (
	"context"
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
	var numOfRequestsSent int
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			r := NewResult()
			res, err := c.Do(req)
			if err != nil {
				break
			}
			if err := res.Body.Close(); err != nil {
				log.Fatalf("failed to close response body: %+v", err)
			}
			numOfRequestsSent++
			r.BuildResult(*res)
			r.printDetails(numOfRequestsSent)
		}
	}
	t.Stop()
}

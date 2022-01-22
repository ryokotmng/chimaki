package chimaki

import (
	"context"
	"log"
	"net/http"
	"time"
)

type client struct {
	endpoint   string
	httpMethod string
	rate       uint64
	*http.Client
}

func NewClient(endpoint string, httpMethod string, rate uint64) *client {
	c := *http.DefaultClient
	clone := &c
	return &client{endpoint, httpMethod, rate, clone}
}

func (c *client) SendRequest(ctx context.Context) {
	req, err := http.NewRequest(c.httpMethod, c.endpoint, nil)
	if err != nil {
		return
	}
	t := time.NewTicker(time.Duration(1000/c.rate) * time.Millisecond)
	var numOfRequestsSent int
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			r := NewResult(c.endpoint)
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

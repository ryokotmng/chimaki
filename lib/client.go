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

type ClientOptions struct {
	Url        string
	HttpMethod string
	Rate       uint64
}

func NewClient(ops *ClientOptions) *client {
	c := *http.DefaultClient
	clone := &c
	return &client{ops.Url, ops.HttpMethod, ops.Rate, clone}
}

func (c *client) ExecuteLoadTest(ctx context.Context) *Results {
	req, err := http.NewRequest(c.httpMethod, c.endpoint, nil)
	if err != nil {
		return nil
	}
	t := time.NewTicker(time.Duration(1000/c.rate) * time.Millisecond)
	var numOfRequestsSent int
	var results Results
Req:
	for {
		select {
		case <-ctx.Done():
			t.Stop()
			break Req
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
			results.Add(r)
		}
	}
	return &results
}

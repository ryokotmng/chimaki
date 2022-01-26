package chimaki

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type client struct {
	endpoint   string
	httpMethod string
	rate       uint64
	body       io.Reader
	*http.Client
}

type ClientOptions struct {
	Url        string
	HttpMethod string
	Rate       uint64
	Bodyf      string
}

func NewClient(ops *ClientOptions) *client {
	c := *http.DefaultClient
	clone := &c
	body, e := ioutil.ReadFile(ops.Bodyf)
	if e != nil {
		body = nil
	}
	var b bytes.Buffer
	b.Write(body)
	return &client{ops.Url, ops.HttpMethod, ops.Rate, &b, clone}
}

func (c *client) ExecuteLoadTest(ctx context.Context) *Results {
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
			req, err := http.NewRequest(c.httpMethod, c.endpoint, c.body)
			if err != nil {
				return nil
			}
			r := NewResult(c.endpoint)
			res, err := c.Do(req)
			numOfRequestsSent++
			if err != nil {
				r.Error = err.Error()
			} else {
				r.BuildWithResponse(*res)
				r.printDetails(numOfRequestsSent)
			}
			results.Add(r)
		}
	}
	return &results
}

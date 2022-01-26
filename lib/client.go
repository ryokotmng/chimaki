package chimaki

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type client struct {
	endpoint   string
	httpMethod string
	rate       uint64
	cpus       int
	body       io.Reader
	*http.Client
}

type ClientOptions struct {
	Url        string
	HttpMethod string
	Rate       uint64
	CPUs       int
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
	return &client{ops.Url, ops.HttpMethod, ops.Rate, ops.CPUs, &b, clone}
}

func (c *client) ExecuteLoadTest(ctx context.Context) *Results {
	var wg sync.WaitGroup
	ch := make(chan *Result)
	t := time.NewTicker(time.Duration(1000/c.rate) * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			t.Stop()
			goto TestFinished
		case <-t.C:
			for i := 0; i < c.cpus; i++ {
				wg.Add(1)
				go c.sendRequest(ch, &wg)
			}
		}
	}
TestFinished:
	var results Results
	go results.receive(ch)
	wg.Wait()
	close(ch)
	return &results
}

func (rs *Results) receive(ch chan *Result) {
	for r := range ch {
		rs.Add(r)
	}
}

func (c *client) sendRequest(ch chan *Result, wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := http.NewRequest(c.httpMethod, c.endpoint, c.body)
	if err != nil {
		return
	}
	r := NewResult(c.endpoint)
	res, err := c.Do(req)
	if err != nil {
		r.Error = err.Error()
	} else {
		r.BuildWithResponse(*res)
		r.printDetails()
	}
	ch <- r
}

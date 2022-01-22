package chimaki

import (
	"fmt"
	"net/http"
	"time"
)

// Result contains the results of a single Target hit.
type Result struct {
	StatusCode uint16        `json:"status_code"`
	Timestamp  time.Time     `json:"timestamp"`
	Latency    time.Duration `json:"latency"`
	Error      string        `json:"error"`
	Body       []byte        `json:"body"`
	HttpMethod string        `json:"http_method"`
	URL        string        `json:"url"`
	Header     http.Header   `json:"header"`
}

func NewResult(url string) *Result {
	return &Result{
		Timestamp: time.Now(),
		URL:       url,
	}
}

func (r *Result) BuildResult(res http.Response) {
	r.StatusCode = uint16(res.StatusCode)
	r.Latency = time.Now().Sub(r.Timestamp)
	r.Header = res.Header
}

// End returns the time at which a Result ended.
func (r *Result) End() time.Time { return r.Timestamp.Add(r.Latency) }

func (r *Result) printDetails(numOfRequestsSent int) {
	fmt.Printf("number of requests sent: %v\n", numOfRequestsSent)
	fmt.Printf("latency: %vms\n", r.Latency)
	fmt.Printf("response(status code: %v) ", r.StatusCode)
}

// Results is a slice of Result type elements.
type Results []Result

// Add implements the Add method of the Report interface by appending the given
// Result to the slice.
func (rs *Results) Add(r *Result) { *rs = append(*rs, *r) }

package chimaki

import (
	"fmt"
	"net/http"
	"sort"
	"time"
)

// Result contains the results of a single Target hit.
type Result struct {
	StatusCode int           `json:"status_code"`
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

func (r *Result) BuildWithResponse(res http.Response) {
	r.StatusCode = res.StatusCode
	r.Latency = time.Now().Sub(r.Timestamp)
	r.Header = res.Header
}

// End returns the time at which a Result ended.
func (r *Result) End() time.Time { return r.Timestamp.Add(r.Latency) }

func (r *Result) printDetails() {
	fmt.Printf("result of single request ---------------\n"+
		"latency: %v\n response(status code: %v)\n",
		r.Latency, r.StatusCode)
}

// Results is a slice of Result type elements.
type Results []Result

// Add implements the Add method of the Report interface by appending the given
// Result to the slice.
func (rs *Results) Add(r *Result) { *rs = append(*rs, *r) }

func (rs *Results) CreateMetrics() {
	var errCount int
	count := len(*rs)
	if count == 0 {
		fmt.Println("test finished! No requests sent")
		return
	}
	m := &metrics{RequestsSent: count, StatusCodes: make(map[int]int)}
	for _, r := range *rs {
		if m.LatencyMetrics.Max < r.Latency {
			m.LatencyMetrics.Max = r.Latency
		}
		if r.Error != "" {
			errCount++
		} else {
			m.StatusCodes[r.StatusCode]++
			m.Latencies = append(m.Latencies, r.Latency)
			m.LatencyMetrics.Total += r.Latency
		}
	}
	m.ErrorRate = float64(errCount) / float64(m.RequestsSent)
	sort.Slice(m.Latencies, func(i, j int) bool { return m.Latencies[i] < m.Latencies[j] })
	m.calcMetrics()
	m.PrintMetrics()
}

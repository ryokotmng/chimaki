package chimaki

import (
	"fmt"
	"net/http"
	"time"
)

type metrics struct {
	Timestamp      time.Time      `json:"timestamp"`
	LatencyMetrics LatencyMetrics `json:"latency_metrics"`
	ErrorRate      float64        `json:"error_rate"`
	Body           []byte         `json:"body"`
	Method         string         `json:"method"`
	URL            string         `json:"url"`
	Header         http.Header    `json:"header"`
	RequestsSent   uint64         `json:"requests_sent"`

	StatusCodes map[uint16]uint64 `json:"status_codes"`
	Errors      map[uint16]uint64 `json:"errors"`
}

func (m *metrics) PrintMetrics() {
	fmt.Println("test finished ==================")
	fmt.Printf("total time of the test: %v\n", m.LatencyMetrics.Total)
	fmt.Printf("latencies: mean %v ms, min %v ms, max %v ms\n",
		m.LatencyMetrics.Mean, m.LatencyMetrics.Min, m.LatencyMetrics.Max)
	fmt.Printf("number of requests sent: %v\n", m.RequestsSent)
	fmt.Printf("mean time of latencies: %vms\n", m.LatencyMetrics.Mean)
}

// LatencyMetrics holds computed request latency metrics.
type LatencyMetrics struct {
	// Total is the total latency sum of all requests in an attack.
	Total time.Duration `json:"total"`
	// Mean is the mean request latency.
	Mean time.Duration `json:"mean"`
	// P50 is the 50th percentile request latency.
	P50 time.Duration `json:"50th"`
	// P90 is the 90th percentile request latency.
	P90 time.Duration `json:"90th"`
	// P95 is the 95th percentile request latency.
	P95 time.Duration `json:"95th"`
	// P99 is the 99th percentile request latency.
	P99 time.Duration `json:"99th"`
	// Max is the maximum observed request latency.
	Max time.Duration `json:"max"`
	// Min is the minimum observed request latency.
	Min time.Duration `json:"min"`
}

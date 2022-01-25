package chimaki

import (
	"fmt"
	"math"
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
	Latencies   []time.Duration   `json:"latencies"`
}

func (m *metrics) calcMetrics() {
	// MEMO: can use goroutine to calc, maybe?
	m.calcMeanValueOfLatencies()
	m.calcLatenciesByPCT()
}

func (m *metrics) calcMeanValueOfLatencies() {
	if m.RequestsSent%2 == 0 {
		idx := m.RequestsSent / 2
		m.LatencyMetrics.Mean = (m.Latencies[idx] + m.Latencies[idx]) / 2
		return
	}
	m.LatencyMetrics.Mean = m.Latencies[(m.RequestsSent-1)/2]
}

func (m *metrics) calcLatenciesByPCT() {
	findLatencyByPCT := func(pct float64) time.Duration {
		numOfRequests := float64(m.RequestsSent)
		return m.Latencies[int(math.Ceil(numOfRequests*pct))]
	}
	m.LatencyMetrics.P50 = findLatencyByPCT(0.50)
	m.LatencyMetrics.P90 = findLatencyByPCT(0.90)
	m.LatencyMetrics.P95 = findLatencyByPCT(0.95)
	m.LatencyMetrics.Min = m.Latencies[0]
	m.LatencyMetrics.Max = m.Latencies[m.RequestsSent-1]
}

func (m *metrics) PrintMetrics() {
	fmt.Println("test finished! overall metrics as below ==================")
	fmt.Printf("total time of the test: %v\n", m.LatencyMetrics.Total)
	fmt.Printf("latencies: mean %v, min %v, max %v\n",
		m.LatencyMetrics.Mean, m.LatencyMetrics.Min, m.LatencyMetrics.Max)
	fmt.Printf("number of requests sent: %v\n", m.RequestsSent)
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

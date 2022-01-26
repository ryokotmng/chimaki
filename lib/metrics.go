package chimaki

import (
	"fmt"
	"math"
	"net/http"
	"sync"
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
	RequestsSent   int            `json:"requests_sent"`

	StatusCodes map[int]int     `json:"status_codes"`
	Latencies   []time.Duration `json:"latencies"`
}

func (m *metrics) calcMetrics() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go m.calcMeanValueOfLatencies(&wg)
	go m.calcLatenciesByPCT(&wg)
	wg.Wait()
}

func (m *metrics) calcMeanValueOfLatencies(wg *sync.WaitGroup) {
	defer wg.Done()
	count := len(m.Latencies)
	if count%2 == 0 {
		idx := count / 2
		m.LatencyMetrics.Mean = (m.Latencies[idx] + m.Latencies[idx]) / 2
		return
	}
	m.LatencyMetrics.Mean = m.Latencies[(count-1)/2]
}

func (m *metrics) calcLatenciesByPCT(wg *sync.WaitGroup) {
	defer wg.Done()
	count := len(m.Latencies)
	findLatencyByPCT := func(pct float64) time.Duration {
		if count == 1 {
			return m.Latencies[0]
		}
		return m.Latencies[int(math.Ceil(float64(count)*pct))]
	}
	m.LatencyMetrics.P50 = findLatencyByPCT(0.50)
	m.LatencyMetrics.P90 = findLatencyByPCT(0.90)
	m.LatencyMetrics.P95 = findLatencyByPCT(0.95)
	m.LatencyMetrics.Min = m.Latencies[0]
	m.LatencyMetrics.Max = m.Latencies[count-1]
}

func (m *metrics) PrintMetrics() {
	fmt.Println("test finished! overall metrics as below ==================")
	fmt.Printf("total time of the test: %v\n", m.LatencyMetrics.Total)
	fmt.Printf("latencies: mean %v, min %v, max %v\n",
		m.LatencyMetrics.Mean, m.LatencyMetrics.Min, m.LatencyMetrics.Max)
	fmt.Printf("number of requests sent: %v\n", m.RequestsSent)
	fmt.Printf("error rate: %v\n", m.ErrorRate)
	fmt.Printf("status codes: %v\n", m.StatusCodes)
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

package metrics

import (
	"time"

	"github.com/0xPolygonHermez/zkevm-node/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	prefix              = "jsonrpc_"
	requestPrefix       = prefix + "request_"
	requestsHandledName = requestPrefix + "handled"
	requestDurationName = requestPrefix + "duration"
	connName            = requestPrefix + "connection"

	requestHandledTypeLabelName = "type"
)

// RequestHandledLabel represents the possible values for the
// `jsonrpc_request_handled` metric `type` label.
type RequestHandledLabel string

// ConnLabel represents the possible values for the
// `jsonrpc_request_connection` metric `type` label.
type ConnLabel string

const (
	// RequestHandledLabelInvalid represents an request of type invalid
	RequestHandledLabelInvalid RequestHandledLabel = "invalid"
	// RequestHandledLabelError represents an request of type error
	RequestHandledLabelError RequestHandledLabel = "error"
	// RequestHandledLabelSingle represents an request of type single
	RequestHandledLabelSingle RequestHandledLabel = "single"
	// RequestHandledLabelBatch represents an request of type batch
	RequestHandledLabelBatch RequestHandledLabel = "batch"

	// HTTPConnLabel represents a HTTP connection
	HTTPConnLabel ConnLabel = "HTTP"
	// WSConnLabel represents a WS connection
	WSConnLabel ConnLabel = "WS"
)

// Register the metrics for the jsonrpc package.
func Register() {
	var (
		counterVecs []metrics.CounterVecOpts
		histograms  []prometheus.HistogramOpts
	)

	counterVecs = []metrics.CounterVecOpts{
		{
			CounterOpts: prometheus.CounterOpts{
				Name: requestsHandledName,
				Help: "[JSONRPC] number of requests handled",
			},
			Labels: []string{requestHandledTypeLabelName},
		},
	}

	start := 0.1
	width := 0.1
	count := 10
	histograms = []prometheus.HistogramOpts{
		{
			Name:    requestDurationName,
			Help:    "[JSONRPC] Histogram for the runtime of requests",
			Buckets: prometheus.LinearBuckets(start, width, count),
		},
	}

	metrics.RegisterCounterVecs(counterVecs...)
	metrics.RegisterHistograms(histograms...)

	// X1 handler
	metrics.RegisterCounterVecs(counterVecsX1...)
	metrics.RegisterHistogramVecs(histogramVecs...)
}

// CountConn increments the connection counter vector by one for the
// given label.
func CountConn(label ConnLabel) {
	metrics.CounterVecInc(connName, string(label))
}

// RequestHandled increments the requests handled counter vector by one for the
// given label.
func RequestHandled(label RequestHandledLabel) {
	metrics.CounterVecInc(requestsHandledName, string(label))
}

// RequestDuration observes (histogram) the duration of a request from the
// provided starting time.
func RequestDuration(start time.Time) {
	metrics.HistogramObserve(requestDurationName, time.Since(start).Seconds())
}

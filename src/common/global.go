package common

import (
	"github.com/kataras/iris/context"
	"github.com/prometheus/client_golang/prometheus"
	"juggernaut/common/grpc"
	"juggernaut/lib/logger"
	"strconv"
	"time"
)

var Logger *logger.Logger
var HttpServerCounterVec *prometheus.CounterVec
var HttpServerTimerVec *prometheus.HistogramVec

func InitLogger(config *logger.Config) (err error) {
	Logger, err = logger.NewLogger(config)
	return err
}

func InitGrpcSrv(config *grpc.Config) {
	grpc.Init(config)
}

func InitPrometheus() {
	constLabels := map[string]string{"service": "iris-gateway", "env": "test", "host": "0.0.0.0"}

	HttpServerTimerVec = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:        "http_request_duration_seconds",
			Help:        "How long it took to process the HTTP request, partitioned by status code, method and HTTP path.",
			ConstLabels: constLabels,
			Buckets:     []float64{0.3, 1.2, 5.0},
		},
		[]string{"code", "method", "path"},
	)

	HttpServerCounterVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name:        "http_requests_total",
			Help:        "How many HTTP requests processed, partitioned by status code, method and HTTP path.",
			ConstLabels: constLabels,
		},
		[]string{"code", "method", "path"},
	)

	prometheus.MustRegister(HttpServerTimerVec, HttpServerCounterVec)
}

func HttpInterceptor(ctx context.Context) {
	start := time.Now()
	ctx.Next()

	r := ctx.Request()
	statusCode := strconv.Itoa(ctx.GetStatusCode())
	duration := float64(time.Since(start).Nanoseconds()) / 1000000000
	labels := []string{statusCode, r.Method, r.URL.Path}

	HttpServerCounterVec.WithLabelValues(labels...).Inc()
	HttpServerTimerVec.WithLabelValues(labels...).Observe(duration)
}

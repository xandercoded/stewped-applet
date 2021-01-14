package metrics

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	"github.com/slok/go-http-metrics/middleware/std"
)

var mw middleware.Middleware

func init() {
	mw = middleware.New(middleware.Config{
		Recorder: prometheus.NewRecorder(prometheus.Config{}),
	})
}

func RegisterHandlerMetrics(router *mux.Router) {
	router.Use(std.HandlerProvider("", mw))
	router.Handle("/metrics", promhttp.Handler())
}

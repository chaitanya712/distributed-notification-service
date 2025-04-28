package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	TotalNotifications = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "notifications_sent_total",
		Help: "Total number of notifications sent",
	})

	FailedNotifications = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "notifications_failed_total",
		Help: "Total number of notifications that failed",
	})
)

func Init() {
	prometheus.MustRegister(TotalNotifications, FailedNotifications)
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":2112", nil)
}
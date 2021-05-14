package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type metrics struct {
	ProcessedCounter prometheus.Counter
	HandledCounter prometheus.Counter
	WelcomeCounter prometheus.Counter
	ConnectSuccessCounter prometheus.Counter
	ConnectFailCounter prometheus.Counter
	KnownUsers prometheus.Gauge
}

func createMetrics() *metrics {
	return &metrics {
		ProcessedCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "irccat_processed_messaged",
			Help: "Number of IRC messages processed",
		}),
		HandledCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "irccat_handled_messaged",
			Help: "Number of IRC messages containing commands",
		}),
		WelcomeCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "irccat_welcome_count",
			Help: "Number of IRC welcomes. This should match connects",
		}),
		ConnectSuccessCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "irccat_connect_success",
			Help: "Number of successful connections",
		}),
		ConnectFailCounter: promauto.NewCounter(prometheus.CounterOpts{
			Name: "irccat_connect_fail",
			Help: "Number of failed connection attempts",
		}),
		KnownUsers: promauto.NewGauge(prometheus.GaugeOpts{
			Name: "irccat_users_gauge",
			Help: "Number of users tracked",
		}),
	}
}

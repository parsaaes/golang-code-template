package metric

import (
	"net/http"

	"github.com/rashadansari/golang-code-template/config"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

func StartPrometheusServer(prometheus config.Prometheus) {
	if prometheus.Enabled {
		metricServer := http.NewServeMux()
		metricServer.Handle("/metrics", promhttp.Handler())

		if err := http.ListenAndServe(prometheus.Address, metricServer); err != nil {
			logrus.Panicf("failed to start prometheus metrics server %s", err.Error())
		}
	}
}

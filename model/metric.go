package model

import (
	"time"

	"github.com/rashadansari/golang-code-template/config"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	LabelRepoName      = "repo_name"
	LabelRepoMethod    = "repo_method"
	ErrorIncrementStep = 1
)

type Metrics struct {
	ErrCounter *prometheus.CounterVec
	Histogram  *prometheus.HistogramVec
}

//nolint:gochecknoglobals
var (
	metrics = Metrics{
		ErrCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: config.Namespace,
				Name:      "repo_error_total",
			}, []string{LabelRepoName, LabelRepoMethod}),

		Histogram: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: config.Namespace,
			Name:      "repo_duration_total",
		}, []string{LabelRepoName, LabelRepoMethod}),
	}

	DoNotReportErrors = []error{}
)

func (m Metrics) report(repoName, methodName string, startTime time.Time, err error) {
	m.Histogram.With(prometheus.Labels{LabelRepoName: repoName, LabelRepoMethod: methodName}).
		Observe(time.Since(startTime).Seconds())

	// Skip do not report errors
	for _, doNotReportError := range DoNotReportErrors {
		if err == doNotReportError {
			return
		}
	}

	if err != nil {
		m.ErrCounter.With(prometheus.Labels{LabelRepoName: repoName, LabelRepoMethod: methodName}).Add(ErrorIncrementStep)
	}
}

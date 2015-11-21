// ----------------------------------------------------------------------------------------------------
// INSTRUMENTATION
// Uses package metrics to record statistics about the service's runtime behavior. Counting the number
// of jobs processed, recording the duration of requests after they've finished, and tracking the number
// of in-flight operations would all be considered instrumentation.
//
// We can use the same middleware pattern that we used for logging.
// ----------------------------------------------------------------------------------------------------

package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

func instrumentingMiddleware(
	requestCount metrics.Counter,
	requestLatency metrics.TimeHistogram,
	countResult metrics.Histogram,
) ServiceMiddleware {
	return func(next NounService) NounService {
		return instrmw{requestCount, requestLatency, countResult, next}
	}
}

type instrmw struct {
	requestCount   metrics.Counter
	requestLatency metrics.TimeHistogram
	countResult    metrics.Histogram
	NounService
}

// Noun
// --------------------------------------------------
func (mw instrmw) Noun(req nounRequest) (output string, err error) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "noun"}
		errorField := metrics.Field{Key: "error", Value: fmt.Sprintf("%v", err)}
		mw.requestCount.With(methodField).With(errorField).Add(1)
		mw.requestLatency.With(methodField).With(errorField).Observe(time.Since(begin))
	}(time.Now())

	output, err = mw.NounService.Place(req)
	return
}

// Place
// --------------------------------------------------
func (mw instrmw) Place(domain string, category string) (output string, err error) {
	defer func(begin time.Time) {
		methodField := metrics.Field{Key: "method", Value: "place"}
		errorField := metrics.Field{Key: "error", Value: fmt.Sprintf("%v", err)}
		mw.requestCount.With(methodField).With(errorField).Add(1)
		mw.requestLatency.With(methodField).With(errorField).Observe(time.Since(begin))
	}(time.Now())

	output, err = mw.NounService.Place(domain, category)
	return
}

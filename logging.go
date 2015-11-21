// ----------------------------------------------------------------------------------------------------
// COMPONENT & APPLICATION LOGGING
//
// COMPONENT LOGGING
// Logging middleware that provides logging at both the application and component levels.
// Logging at the application level allows us to log things like the parameters passed into the service.
// Logging at the component level provides logging to transport-level data.
//
// APPLICATION LOGGING
// Defines middleware for the service that provides logging at the application domain level.
// Alows for logging of the parameters that are passed in.
// Since our Service is defined as an interface, we just need to make a new type which wraps an existing Service,
// and performs the extra logging duties.
//
// Logging Reference
// ----------------------------------------------------------------------------------------------------
// INFO LEVEL
// The start and end of a request
// Successful Auth attempt
//
// DEBUG LEVEL
// Any parameters passed to an endpoint
//
// ERROR LEVEL
// Handled exceptions
// Invalid auth attempts
//
// FATAL LEVEL
// Unhandled exceptions
//
// ----------------------------------------------------------------------------------------------------

package main

import (
	"time"

	"github.com/go-kit/kit/log"
)

// Basic logging middleware.
// --------------------------------------------------
func loggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next NounService) NounService {
		return logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	NounService
}

// Logging for the Noun method
// --------------------------------------------------
func (mw logmw) Noun(req nounRequest) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "noun",
			"input", req,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.NounService.Noun(req)
	return
}

// Logging for the Place method
// --------------------------------------------------
func (mw logmw) Place(domain string, category string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "place",
			"input", domain,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.NounService.Place(domain, category)
	return
}

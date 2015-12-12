// ==================================================
// Nouns
// ==================================================
// Start in proxy mode
// nouns -listen=:8001 &
// nouns -listen=:8002 &
// nouns -listen=:8003 &
// nouns -listen=:9000 -proxy=localhost:8001,localhost:8002,localhost:8003
//
// Or Run all at once?
// nouns -listen=:8001 & nouns -listen=:8002 & nouns -listen=:8003 & nouns -listen=:9000 -proxy=localhost:8001,localhost:8002,localhost:8003
//
// Posting to the nouns service
// curl -XPOST -d '{"s":"hello, world"}' localhost:9000/place
// curl -X POST -d "$(cat ~/Desktop/place.json)" localhost:9000/place
//
// Kill the servers
// kill -9  $(ps aux | grep listen | grep -v grep | awk '{print $2}')
// ==================================================

package main

import (
	"flag"
	"net/http"
	"os"
	"time"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	// Fire up a simple HTTP server
	var (
		listen = flag.String("listen", ":9000", "HTTP listen address")
		proxy  = flag.String("proxy", "", "Optional comma-separated list of URLs to proxy requests")
	)
	flag.Parse()

	// Initialize Logging
	// --------------------------------------------------
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.NewContext(logger).With("listen", *listen).With("caller", log.DefaultCaller)

	ctx := context.Background()

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounter(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := metrics.NewTimeHistogram(time.Microsecond, kitprometheus.NewSummary(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys))
	countResult := kitprometheus.NewSummary(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})

	// Initialize Service
	// --------------------------------------------------
	var svc NounService
	svc = nounService{}
	svc = proxyingMiddleware(*proxy, ctx, logger)(svc)
	svc = loggingMiddleware(logger)(svc)
	svc = instrumentingMiddleware(requestCount, requestLatency, countResult)(svc)

	// Noun Handler
	// --------------------------------------------------
	nounHandler := httptransport.NewServer(
		ctx,
		makeNounEndpoint(svc),
		decodeNounRequest,
		encodeResponse,
	)

	// Place Handler
	// --------------------------------------------------
	placeHandler := httptransport.NewServer(
		ctx,
		makePlaceEndpoint(svc),
		decodePlaceRequest,
		encodeResponse,
	)

	// Define Endpoints
	// --------------------------------------------------
	http.Handle("/noun", nounHandler)
	http.Handle("/place", placeHandler)
	http.Handle("/metrics", stdprometheus.Handler())
	_ = logger.Log("msg", "HTTP", "addr", *listen)
	_ = logger.Log("err", http.ListenAndServe(*listen, nil))
}

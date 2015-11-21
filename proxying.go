// ----------------------------------------------------------------------------------------------------
// SERVICE PROXY
// Proxying middleware to facilitate calling other services.
// ----------------------------------------------------------------------------------------------------

package main

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"

	jujuratelimit "github.com/juju/ratelimit"
	"github.com/sony/gobreaker"
	"golang.org/x/net/context"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/loadbalancer"
	"github.com/go-kit/kit/loadbalancer/static"
	"github.com/go-kit/kit/log"
	kitratelimit "github.com/go-kit/kit/ratelimit"
	httptransport "github.com/go-kit/kit/transport/http"
)

func proxyingMiddleware(proxyList string, ctx context.Context, logger log.Logger) ServiceMiddleware {
	if proxyList == "" {
		_ = logger.Log("proxy_to", "none")
		return func(next NounService) NounService { return next }
	}
	proxies := split(proxyList)
	_ = logger.Log("proxy_to", fmt.Sprint(proxies))

	return func(next NounService) NounService {
		var (
			qps         = 100 // max to each instance
			publisher   = static.NewPublisher(proxies, factory(ctx, qps), logger)
			lb          = loadbalancer.NewRoundRobin(publisher)
			maxAttempts = 3
			maxTime     = 100 * time.Millisecond
			endpoint    = loadbalancer.Retry(maxAttempts, maxTime, lb)
		)
		return proxymw{ctx, endpoint, next}
	}
}

// proxymw implements NounService, forwarding endpoint requests to the
// provided endpoint, and serving all other (i.e. Count) requests via the
// embedded NounService.
type proxymw struct {
	context.Context
	NounEndpoint endpoint.Endpoint
	// UppercaseEndpoint endpoint.Endpoint
	NounService
}

// Noun Proxy
// --------------------------------------------------
func (mw proxymw) Noun(req nounRequest) (string, error) {
	response, err := mw.NounEndpoint(mw.Context, nounRequest{Req: req})
	if err != nil {
		return "", err
	}

	resp := response.(nounResponse)
	if resp.Err != "" {
		return resp.V, errors.New(resp.Err)
	}
	return resp.V, nil
}

func factory(ctx context.Context, qps int) loadbalancer.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		var e endpoint.Endpoint
		e = makeNounProxy(ctx, instance)
		e = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(e)
		e = kitratelimit.NewTokenBucketLimiter(jujuratelimit.NewBucketWithRate(float64(qps), int64(qps)))(e)
		return e, nil, nil
	}
}

func makeNounProxy(ctx context.Context, instance string) endpoint.Endpoint {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		panic(err)
	}
	if u.Path == "" {
		u.Path = "/noun"
	}
	return httptransport.NewClient(
		"GET",
		u,
		encodeRequest,
		decodeNounResponse,
	).Endpoint()
}

// --------------------------------------------------
func split(s string) []string {
	a := strings.Split(s, ",")
	for i := range a {
		a[i] = strings.TrimSpace(a[i])
	}
	return a
}

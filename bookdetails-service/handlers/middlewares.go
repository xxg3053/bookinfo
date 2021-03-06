package handlers

import (
	"bookinfo/bookdetails-service/svc"
	pb "bookinfo/pb/details"
	"time"
	"github.com/go-kit/kit/ratelimit"
	"golang.org/x/time/rate"
	"github.com/go-kit/kit/circuitbreaker"
	"github.com/sony/gobreaker"
)

// WrapEndpoints accepts the service's entire collection of endpoints, so that a
// set of middlewares can be wrapped around every middleware (e.g., access
// logging and instrumentation), and others wrapped selectively around some
// endpoints and not others (e.g., endpoints requiring authenticated access).
// Note that the final middleware wrapped will be the outermost middleware
// (i.e. applied first)
func WrapEndpoints(in svc.Endpoints) svc.Endpoints {

	// Pass a middleware you want applied to every endpoint.
	// optionally pass in endpoints by name that you want to be excluded
	// e.g.
	// in.WrapAllExcept(authMiddleware, "Status", "Ping")

	// Pass in a svc.LabeledMiddleware you want applied to every endpoint.
	// These middlewares get passed the endpoints name as their first argument when applied.
	// This can be used to write generic metric gathering middlewares that can
	// report the endpoint name for free.
	// github.com/tuneinc/truss/_example/middlewares/labeledmiddlewares.go for examples.
	// in.WrapAllLabeledExcept(errorCounter(statsdCounter), "Status", "Ping")

	// How to apply a middleware to a single endpoint.
	// in.ExampleEndpoint = authMiddleware(in.ExampleEndpoint)

	//in.DetailEndpoint =

	//日志
	in.DetailEndpoint = LoggingEndpointMiddleware()(in.DetailEndpoint)

	//全链路追踪
	in.DetailEndpoint = ZipkinEndpointMiddleware()(in.DetailEndpoint)

	//限频
	in.DetailEndpoint = ratelimit.NewErroringLimiter(rate.NewLimiter(
		1000, //每秒产生的令牌数
		10000, //令牌池最大值
	))(in.DetailEndpoint)

	//熔断
	in.DetailEndpoint = circuitbreaker.Gobreaker(
		gobreaker.NewCircuitBreaker(gobreaker.Settings{
			Name:        "/v1/detail",
			MaxRequests: 50000,
			Interval:    1 * time.Second,
		}))(in.DetailEndpoint)

	return in
}

func WrapService(in pb.BookDetailsServer) pb.BookDetailsServer {

	//日志
	in = LoggingSvcMiddleware()(in)

	//prometheus 中间件
	in = Instrumenting()(in)

	return in
}

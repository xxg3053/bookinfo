// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 7dc4d5d85c
// Version Date: 2018年 5月28日 星期一 22时12分59秒 UTC

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	pb "bookinfo/pb/comments"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	PostEndpoint endpoint.Endpoint
	GetEndpoint  endpoint.Endpoint
}

// Endpoints

func (e Endpoints) Post(ctx context.Context, in *pb.PostReq) (*pb.PostResp, error) {
	response, err := e.PostEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.PostResp), nil
}

func (e Endpoints) Get(ctx context.Context, in *pb.GetReq) (*pb.GetResp, error) {
	response, err := e.GetEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.GetResp), nil
}

// Make Endpoints

func MakePostEndpoint(s pb.BookCommentsServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.PostReq)
		v, err := s.Post(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeGetEndpoint(s pb.BookCommentsServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.GetReq)
		v, err := s.Get(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAllExcept wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllExcept(middleware, "Status", "Ping")
func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"Post": struct{}{},
		"Get":  struct{}{},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc, _ := range included {
		if inc == "Post" {
			e.PostEndpoint = middleware(e.PostEndpoint)
		}
		if inc == "Get" {
			e.GetEndpoint = middleware(e.GetEndpoint)
		}
	}
}

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(string, endpoint.Endpoint) endpoint.Endpoint

// WrapAllLabeledExcept wraps each Endpoint field of struct Endpoints with a
// LabeledMiddleware, which will receive the name of the endpoint. See
// LabeldMiddleware. See method WrapAllExept for details on excluded
// functionality.
func (e *Endpoints) WrapAllLabeledExcept(middleware func(string, endpoint.Endpoint) endpoint.Endpoint, excluded ...string) {
	included := map[string]struct{}{
		"Post": struct{}{},
		"Get":  struct{}{},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc, _ := range included {
		if inc == "Post" {
			e.PostEndpoint = middleware("Post", e.PostEndpoint)
		}
		if inc == "Get" {
			e.GetEndpoint = middleware("Get", e.GetEndpoint)
		}
	}
}

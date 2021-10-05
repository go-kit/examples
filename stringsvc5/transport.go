package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func MakeHTTPHandler(svc StringService) http.Handler {

	r := gin.Default()

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)
	r.POST("/uppercase", gin.WrapH(uppercaseHandler))

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)
	r.POST("/count", gin.WrapH(countHandler))

	return r
}

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}

func decodeUppercaseRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

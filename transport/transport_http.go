package transport

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/ong-gtp/choice-service/endpoint"
	"github.com/ong-gtp/choice-service/helper"
	"github.com/ong-gtp/choice-service/service"
)

// NewHttpServer defines the http handler endpoints for choice service
func NewHttpServer(svc service.Service, logger log.Logger) *mux.Router {
	// options provided by the go-kit to facilitate error control
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(helper.EncodeErrorResponse),
	}

	// define health endpoint handler
	healthHandler := httptransport.NewServer(
		endpoint.MakeGetHealthEndpoint(svc), //use the endpoint
		helper.DecodeEmptyRequest,           //converts the parameters received via the request body into the struct expected by the endpoint
		helper.EncodeHttpResponse,           //converts the struct returned by the endpoint to a json response
		options...,
	)

	// define choice endpoint handler
	choiceHandler := httptransport.NewServer(
		endpoint.MakeGetRandomChoiceEndpoint(svc), //use the endpoint
		helper.DecodeEmptyRequest,                 //converts the parameters received via the request body into the struct expected by the endpoint
		helper.EncodeHttpResponse,                 //converts the struct returned by the endpoint to a json response
		options...,
	)

	// define choices endpoint  handler
	choicesHandler := httptransport.NewServer(
		endpoint.MakeGetChoicesEndpoint(svc), //use the endpoint
		helper.DecodeEmptyRequest,            //converts the parameters received via the request body into the struct expected by the endpoint
		helper.EncodeHttpResponse,            //converts the struct returned by the endpoint to a json response
		options...,
	)

	r := mux.NewRouter()
	r.Methods("GET").Path("/choice").Handler(choiceHandler)
	r.Methods("GET").Path("/choices").Handler(choicesHandler)
	r.Methods("GET").Path("/choicesv/health").Handler(healthHandler)
	return r
}

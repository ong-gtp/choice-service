package helper

import (
	"context"
	"encoding/json"
	"net/http"
)

// ChoicesSize is the count of choices available
const ChoicesSize = 5

// RandNumServiceUrl is the url to the external random number service
const RandNumServiceUrl = "https://codechallenge.boohma.com/random"

// EncodeHttpResponse converts the struct returned by the endpoint to a json response
func EncodeHttpResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

// DecodeEmptyRequest converts the parameters received via the request body into the struct expected by the endpoint
func DecodeEmptyRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request interface{}
	return request, nil
}

// EncodeErrorResponse converts errors to json response
func EncodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

// codeFrom returns the http status code from service errors
func codeFrom(err error) int {
	switch err {
	case ErrRandomNumberService:
		return http.StatusBadGateway
	case ErrRandomNumberBody:
		return http.StatusBadGateway
	case ErrRandomNumberUnmarshal:
		return http.StatusBadGateway
	case ErrRandomChoice:
		return http.StatusNoContent
	case ErrRandomNumberValue:
		return http.StatusNoContent
	default:
		return http.StatusInternalServerError
	}
}

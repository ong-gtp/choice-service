package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/ong-gtp/choice-service/service"
)

// MakeHealthEndpoint creates the go-kit enpoint for GetHealth
func MakeGetHealthEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		cs, err := svc.GetHealth(ctx)
		if err != nil {
			return "", err
		}
		return cs, nil
	}
}

// MakeChoicesEndpoint creates the go-kit enpoint for GetChoices
func MakeGetChoicesEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		cs, err := svc.GetChoices(ctx)
		if err != nil {
			return "", err
		}
		return cs, nil
	}
}

// MakeChoiceEndpoint creates the go-kit enpoint for GetChoice
func MakeGetRandomChoiceEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		c, err := svc.GetRandomChoice(ctx)
		if err != nil {
			return "", err
		}
		return c, nil
	}
}

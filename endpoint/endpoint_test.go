package endpoint

import (
	"context"
	"testing"

	"github.com/ong-gtp/choice-service/service"
	"github.com/stretchr/testify/assert"
)

func TestMakeGetHealthEndpoint(t *testing.T) {
	s := service.NewService()
	ep := MakeGetHealthEndpoint(s)
	t.Run("health", func(t *testing.T) {
		var request interface{}
		r, err := ep(context.Background(), request)
		if err != nil {
			t.Errorf("expected %v, got %v", nil, err)
		}
		assert.Equal(t, "ok", r)
	})
}

func TestMakeChoiceEndpoint(t *testing.T) {
	s := service.NewService()
	ep := MakeGetRandomChoiceEndpoint(s)
	t.Run("make choice endoint is not breaking", func(t *testing.T) {
		var request interface{}
		_, err := ep(context.Background(), request)
		if err != nil {
			t.Errorf("expected %v, got %v", nil, err)
		}
	})
}

func TestMakeChoicesEndpoint(t *testing.T) {
	s := service.NewService()
	ep := MakeGetChoicesEndpoint(s)

	t.Run("make choices endpoint is not breaking", func(t *testing.T) {
		var request interface{}
		_, err := ep(context.Background(), request)

		if err != nil {
			t.Errorf("expected %v, got %v", nil, err)
		}
	})
}

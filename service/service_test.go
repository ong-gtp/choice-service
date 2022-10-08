package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChoiceService(t *testing.T) {
	service := NewService()

	t.Run("health", func(t *testing.T) {
		c, _ := service.GetHealth(context.Background())
		assert.Equal(t, "ok", c)
	})

	t.Run("list choices", func(t *testing.T) {
		c, _ := service.GetChoices(context.Background())
		assert.Equal(t, Choices, c)
	})

	t.Run("random choice", func(t *testing.T) {
		c, _ := service.GetRandomChoice(context.Background())
		assert.Contains(t, Choices, c)
	})
}

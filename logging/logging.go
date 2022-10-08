package logging

import (
	"context"
	"time"

	"github.com/go-kit/log"
	"github.com/ong-gtp/choice-service/entities"
	"github.com/ong-gtp/choice-service/service"
)

// NewLoggingMiddleware returns the logging middleware for the service
func NewLoggingMiddleware(logger log.Logger, next service.Service) logmw {
	return logmw{logger, next}
}

// logmw defines the interface for the logging middleware implemented for service.Service
type logmw struct {
	logger log.Logger
	service.Service
}

// GetHealth defines the logging middleware for service.GetHealth
func (mw logmw) GetHealth(ctx context.Context) (response string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "getHealth",
			"input", "",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	response, err = mw.Service.GetHealth(ctx)
	return
}

// GetRandomChoice defines the logging middleware for service.GetRandomChoice
func (mw logmw) GetRandomChoice(ctx context.Context) (choice entities.Choice, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "getRandomChoice",
			"input", "",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	choice, err = mw.Service.GetRandomChoice(ctx)
	return
}

// GetChoices defines the logging middleware for service.GetChoices
func (mw logmw) GetChoices(ctx context.Context) (choices []entities.Choice, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "getChoices",
			"input", "",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	choices, err = mw.Service.GetChoices(ctx)
	return
}

package helper

import "errors"

// service errors
var (
	ErrRandomNumberService   = errors.New("random-number service failed")
	ErrRandomChoice          = errors.New("failed to choose a random choice")
	ErrRandomNumberUnmarshal = errors.New("failed to unmarchal random number service response")
	ErrRandomNumberBody      = errors.New("error in random-number's service's response body")
	ErrRandomNumberValue     = errors.New("error in random-number's value")
)

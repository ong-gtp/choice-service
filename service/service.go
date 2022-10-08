package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ong-gtp/choice-service/entities"
	"github.com/ong-gtp/choice-service/helper"
)

// Service defines the interface for choice-service
type Service interface {
	GetHealth(ctx context.Context) (string, error)
	GetChoices(ctx context.Context) ([]entities.Choice, error)
	GetRandomChoice(ctx context.Context) (entities.Choice, error)
}

// service is the struct that will implement Service interfacce
type service struct{}

// NewService creates a new instance of service
func NewService() *service {
	return &service{}
}

// Choices is the list of choices that can be sent to client
var Choices = []entities.Choice{
	{Id: 1, Name: "Rock"},
	{Id: 2, Name: "Paper"},
	{Id: 3, Name: "Scissors"},
	{Id: 4, Name: "Spook"},
	{Id: 5, Name: "Lizard"},
}

// GetChoices returns all the choices that can be sent to client
func (s *service) GetHealth(ctx context.Context) (string, error) {
	return "ok", nil
}

// GetChoices returns all the choices that can be sent to client
func (s *service) GetChoices(ctx context.Context) ([]entities.Choice, error) {
	return Choices, nil
}

// GetRandomChoice calls an external random number generator and uses the value received
// for evaluating the choice to send to client
func (s *service) GetRandomChoice(ctx context.Context) (entities.Choice, error) {
	response, err := http.Get(helper.RandNumServiceUrl)
	if err != nil {
		return entities.Choice{}, helper.ErrRandomNumberService
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return entities.Choice{}, helper.ErrRandomNumberBody
	}

	var randData entities.RandomNumber
	err = json.Unmarshal(responseData, &randData)
	if err != nil {
		return entities.Choice{}, helper.ErrRandomNumberUnmarshal
	}

	if randData.RandomNumber < 1 || randData.RandomNumber > 100 {
		return entities.Choice{}, helper.ErrRandomNumberValue
	}

	id := (randData.RandomNumber % helper.ChoicesSize) + 1
	for _, val := range Choices {
		if val.Id == id {
			return val, nil
		}
	}

	return entities.Choice{}, helper.ErrRandomChoice
}

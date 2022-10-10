package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ong-gtp/choice-service/entities"
	"github.com/ong-gtp/choice-service/helper"
	"github.com/ong-gtp/choice-service/repository"
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

var repo repository.Repository = repository.NewRepository()

// Choices is the list of choices that can be sent to client
var Choices = repo.GetChoices()

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
		helper.Log("error", "error :", err)
		return entities.Choice{}, helper.ErrRandomNumberService
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		helper.Log("error", "error :", err)
		return entities.Choice{}, helper.ErrRandomNumberBody
	}

	var randData entities.RandomNumber
	err = json.Unmarshal(responseData, &randData)
	if err != nil {
		helper.Log("error", "error :", err)
		return entities.Choice{}, helper.ErrRandomNumberUnmarshal
	}

	if randData.RandomNumber < 1 || randData.RandomNumber > 100 {
		return entities.Choice{}, helper.ErrRandomNumberValue
	}

	helper.Log("debug", "RandomNumber: ", randData.RandomNumber)

	id := (randData.RandomNumber % helper.ChoicesSize) + 1
	for _, val := range Choices {
		if val.Id == id {
			return val, nil
		}
	}

	return entities.Choice{}, helper.ErrRandomChoice
}

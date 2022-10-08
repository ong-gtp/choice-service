package repository

import "github.com/ong-gtp/choice-service/entities"

type Repository interface {
	GetChoices() []entities.Choice
}

type repository struct{}

func NewRepository() *repository {
	return &repository{}
}

var choices = []entities.Choice{
	{Id: 1, Name: "rock"},
	{Id: 2, Name: "paper"},
	{Id: 3, Name: "scissors"},
	{Id: 4, Name: "lizard"},
	{Id: 5, Name: "spock"},
}

func (r *repository) GetChoices() []entities.Choice {
	return choices
}

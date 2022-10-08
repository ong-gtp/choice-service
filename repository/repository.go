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
	{Id: 1, Name: "Rock"},
	{Id: 2, Name: "Paper"},
	{Id: 3, Name: "Scissors"},
	{Id: 4, Name: "Spook"},
	{Id: 5, Name: "Lizard"},
}

func (r *repository) GetChoices() []entities.Choice {
	return choices
}

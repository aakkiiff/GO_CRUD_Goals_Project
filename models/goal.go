package models

import "github.com/kamva/mgm/v3"

type Goal struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Description      string `json:"description" bson:"description"`
}

func NewGoal(name string, description string) *Goal {
	return &Goal{
		Name:        name,
		Description: description,
	}
}

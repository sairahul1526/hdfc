package model

import "github.com/kamva/mgm/v3"

type Category struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Type             string `json:"type" bson:"type"`
}

package model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	mgm.DefaultModel `bson:",inline"`
	CategoryID       primitive.ObjectID `json:"category_id" bson:"category_id"`
	Name             string             `json:"name" bson:"name"`
	Price            float64            `json:"price" bson:"price"`
	Count            int32              `json:"count" bson:"count"`
}

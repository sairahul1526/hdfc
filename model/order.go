package model

import (
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	mgm.DefaultModel `bson:",inline"`
	CustomerID       primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	TotalAmount      float64            `json:"total_amount" bson:"total_amount"`
	Discount         float64            `json:"discount" bson:"discount"`
	FinalAmount      float64            `json:"final_amount" bson:"final_amount"`
	Status           string             `json:"status" bson:"status"`
	DispatchDate     time.Time          `json:"dispatch_date" bson:"dispatch_date"`
	Items            []OrderItem        `json:"items" bson:"items"`
}

type OrderItem struct {
	ProductID   primitive.ObjectID `json:"product_id" bson:"product_id"`
	Amount      float64            `json:"amount" bson:"amount"`
	Count       int8               `json:"count" bson:"count"`
	TotalAmount float64            `json:"total_amount" bson:"total_amount"`
}

package product

import (
	CONSTANT "hdfc-backend/constant"
	MODEL "hdfc-backend/model"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// get all orders by customer id, page
func getOrdersByCustomerID(customerID string, page int) (orders []MODEL.Order, err error) {
	custID, _ := primitive.ObjectIDFromHex(customerID)
	err = mgm.Coll(&MODEL.Order{}).SimpleFind(&orders, bson.M{
		"customer_id": custID,
	},
		options.Find().
			SetLimit(CONSTANT.NumberOfResultsPerPage).
			SetSkip(int64(page)*CONSTANT.NumberOfResultsPerPage),
	)
	return
}

// get order by id
func getOrderByID(orderID string) (order MODEL.Order, err error) {
	id, _ := primitive.ObjectIDFromHex(orderID)
	err = mgm.Coll(&MODEL.Order{}).FindByID(id, &order)
	return
}

// get all products with multiple ids
func getProductsByIDs(productIDs []string) (products []MODEL.Product, err error) {
	ids := make([]primitive.ObjectID, len(productIDs))
	for i, id := range productIDs {
		ids[i], _ = primitive.ObjectIDFromHex(id)
	}
	err = mgm.Coll(&MODEL.Product{}).SimpleFind(&products, bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	})
	return
}

// get all categories with multiple ids
func getCategoriesByIDs(categoryIDs []string) (categories []MODEL.Category, err error) {
	ids := make([]primitive.ObjectID, len(categoryIDs))
	for i, id := range categoryIDs {
		ids[i], _ = primitive.ObjectIDFromHex(id)
	}
	err = mgm.Coll(&MODEL.Category{}).SimpleFind(&categories, bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	})
	return
}

// decrease product count
func decreaseProductQuantity(product MODEL.Product, count int32) (err error) {
	product.Count -= count
	err = mgm.Coll(&MODEL.Product{}).Update(&product)
	return
}

// update order
func updateOrder(order MODEL.Order) (err error) {
	err = mgm.Coll(&MODEL.Order{}).Update(&order)
	return
}

package product

import (
	CONSTANT "hdfc-backend/constant"
	MODEL "hdfc-backend/model"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// get all categories
func getAllCategories() (categories []MODEL.Category, err error) {
	err = mgm.Coll(&MODEL.Category{}).SimpleFind(&categories, bson.M{})
	return
}

// get all products by category id, page
func getProductsByCategoryID(categoryID string, page int) (products []MODEL.Product, err error) {
	catID, _ := primitive.ObjectIDFromHex(categoryID)
	err = mgm.Coll(&MODEL.Product{}).SimpleFind(&products, bson.M{
		"category_id": catID,
	},
		options.Find().
			SetLimit(CONSTANT.NumberOfResultsPerPage).
			SetSkip(int64(page)*CONSTANT.NumberOfResultsPerPage),
	)
	return
}

// get product by id
func getProductByID(productID string) (product MODEL.Product, err error) {
	err = mgm.Coll(&MODEL.Product{}).FindByID(productID, &product)
	return
}

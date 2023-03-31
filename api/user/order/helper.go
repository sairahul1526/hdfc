package product

import (
	"errors"
	CONSTANT "hdfc-backend/constant"
	MODEL "hdfc-backend/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// check if all products are available
func checkIfProductsAreAvailable(products []MODEL.Product, body OrderCreateRequestBody) (err error) {
	// convert products into map of product id and product
	productsMap := make(map[string]MODEL.Product)
	for _, product := range products {
		productsMap[product.ID.Hex()] = product
	}

	// check if max quantity of a particular product is 10
	for _, product := range body.Products {
		if product.Quantity > 10 {
			return CONSTANT.ProductMaxQuantityError
		}
	}

	// check if all products are available
	for _, product := range body.Products {
		if productsMap[product.ProductID].Count < int32(product.Quantity) {
			return errors.New(productsMap[product.ProductID].Name + " is not available")
		}
	}

	return nil
}

// get product ids from request body
func getProductIDs(body OrderCreateRequestBody) (productIDs []string) {
	for _, product := range body.Products {
		productIDs = append(productIDs, product.ProductID)
	}
	return
}

// build order object
func buildOrder(products []MODEL.Product, body OrderCreateRequestBody) (order MODEL.Order) {
	custID, _ := primitive.ObjectIDFromHex(body.CustomerID)
	order.CustomerID = custID

	order.Status = CONSTANT.OrderStatusPlaced

	// convert products into map of product id and product
	productsMap := make(map[string]MODEL.Product)
	for _, product := range products {
		productsMap[product.ID.Hex()] = product
	}

	// fill total amount for each product like quantity * price
	for _, product := range body.Products {
		productID, _ := primitive.ObjectIDFromHex(product.ProductID)
		order.Items = append(order.Items, MODEL.OrderItem{
			ProductID:   productID,
			Amount:      productsMap[product.ProductID].Price,
			Count:       product.Quantity,
			TotalAmount: productsMap[product.ProductID].Price * float64(product.Quantity),
		})
	}

	// fill total amount for order like sum of all products total amount
	for _, item := range order.Items {
		order.TotalAmount += item.TotalAmount
	}

	order.FinalAmount = order.TotalAmount

	return
}

// check for discounts
func applyDiscounts(order *MODEL.Order, products []MODEL.Product) {

	// convert products into map of product id and product
	productsMap := make(map[string]MODEL.Product)
	for _, product := range products {
		productsMap[product.ID.Hex()] = product
	}

	// get all category ids from products
	var categoryIDs []string
	for _, product := range products {
		categoryIDs = append(categoryIDs, product.CategoryID.Hex())
	}

	// get all categories with multiple ids
	categories, _ := getCategoriesByIDs(categoryIDs)

	// convert categories into map of category id and category
	categoriesMap := make(map[string]MODEL.Category)
	for _, category := range categories {
		categoriesMap[category.ID.Hex()] = category
	}

	// count premium products
	premiumCount := 0
	for _, item := range order.Items {
		if categoriesMap[productsMap[item.ProductID.Hex()].CategoryID.Hex()].Type == CONSTANT.CategoryTypePremium {
			premiumCount++
		}
	}

	// apply 10% discount if premium products are more than 3
	if premiumCount >= 3 {
		order.Discount = order.TotalAmount * 0.1
		order.FinalAmount -= order.FinalAmount * 0.1
	}
}

// update product count after order is placed
func updateProductCount(products []MODEL.Product, body OrderCreateRequestBody) (err error) {
	// convert products into map of product id and product
	productsMap := make(map[string]MODEL.Product)
	for _, product := range products {
		productsMap[product.ID.Hex()] = product
	}

	// update product count
	for _, product := range body.Products {
		decreaseProductQuantity(productsMap[product.ProductID], int32(product.Quantity))
	}

	return
}

package product

// OrderCreateRequestBody - request body for creating new order
type OrderCreateRequestBody struct {
	CustomerID string `json:"customer_id"`
	Products   []struct {
		ProductID string `json:"product_id"`
		Quantity  int8   `json:"quantity"`
	} `json:"products"`
}

// OrderUpdateRequestBody - request body for updating order
type OrderUpdateRequestBody struct {
	Status string `json:"status"`
}

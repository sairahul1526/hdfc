package product

import (
	CONSTANT "hdfc-backend/constant"
	MODEL "hdfc-backend/model"
	UTIL "hdfc-backend/util"
	"net/http"
	"strconv"

	"github.com/kamva/mgm/v3"
)

// OrdersGet - get all orders by customer id
//
// @Summary Get all category details
// @Param customer_id query string false "customer id"
// @Tags Order
// @Router /orders [get]
func OrdersGet(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	page, _ := strconv.Atoi(r.FormValue("page"))
	if page > 0 {
		page -= 1
	}

	// get orders by customer id
	orders, err := getOrdersByCustomerID(r.FormValue("customer_id"), page)
	if err != nil {
		UTIL.SetReponse(w, CONSTANT.StatusCodeServerError, "", CONSTANT.ShowDialog, err.Error(), response)
		return
	}

	response["orders"] = orders
	response["page"] = page + 1

	UTIL.SetReponse(w, CONSTANT.StatusCodeOk, "", CONSTANT.NoDialog, "", response)
}

// OrderPost - create new order with multiple products in post body
//
// @Summary Create new order
// @Param order body OrderCreateRequestBody true "order details"
// @Tags Order
// @Router /order [post]
func OrderPost(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	body := OrderCreateRequestBody{}

	// read request body
	err := UTIL.ReadRequestBody(r, &body)
	if err != nil {
		UTIL.SetReponse(w, CONSTANT.StatusCodeBadRequest, "", CONSTANT.ShowDialog, err.Error(), response)
		return
	}

	// validate customer id

	// get products by ids
	products, err := getProductsByIDs(getProductIDs(body))
	if err != nil {
		UTIL.SetReponse(w, CONSTANT.StatusCodeServerError, "", CONSTANT.ShowDialog, err.Error(), response)
		return
	}

	// check if all products are available
	if err = checkIfProductsAreAvailable(products, body); err != nil {
		UTIL.SetReponse(w, CONSTANT.StatusCodeBadRequest, "", CONSTANT.ShowDialog, err.Error(), response)
		return
	}

	// build order
	order := buildOrder(products, body)

	// apply discount if any
	applyDiscounts(&order, products)

	err = mgm.Coll(&MODEL.Order{}).Create(&order)
	if err != nil {
		UTIL.SetReponse(w, CONSTANT.StatusCodeServerError, "", CONSTANT.ShowDialog, err.Error(), response)
		return
	}

	// update product quantity
	updateProductCount(products, body)

	response["order"] = order

	UTIL.SetReponse(w, CONSTANT.StatusCodeOk, "", CONSTANT.NoDialog, "", response)
}

// OrderUpdate - create new order with multiple products in post body
//
// @Summary Create new order
// @Param id query string true "order id"
// @Param order body OrderUpdateRequestBody true "order details"
// @Tags Order
// @Router /order [patch]
func OrderUpdate(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	body := OrderUpdateRequestBody{}

	// read request body
	err := UTIL.ReadRequestBody(r, &body)
	if err != nil {
		UTIL.SetReponse(w, CONSTANT.StatusCodeBadRequest, "", CONSTANT.ShowDialog, err.Error(), response)
		return
	}

	// get order by id
	order, err := getOrderByID(r.FormValue("id"))
	if err != nil {
		UTIL.SetReponse(w, CONSTANT.StatusCodeServerError, "", CONSTANT.ShowDialog, err.Error(), response)
		return
	}

	// validate order status
	switch order.Status {
	case CONSTANT.OrderStatusPlaced:
	case CONSTANT.OrderStatusDispatched:
		if body.Status == CONSTANT.OrderStatusPlaced {
			UTIL.SetReponse(w, CONSTANT.StatusCodeBadRequest, "", CONSTANT.ShowDialog, CONSTANT.InvalidOrderStatusMessage, response)
			return
		}
	case CONSTANT.OrderStatusCompleted:
		UTIL.SetReponse(w, CONSTANT.StatusCodeBadRequest, "", CONSTANT.ShowDialog, CONSTANT.InvalidOrderStatusMessage, response)
		return
	case CONSTANT.OrderStatusCancelled:
		UTIL.SetReponse(w, CONSTANT.StatusCodeBadRequest, "", CONSTANT.ShowDialog, CONSTANT.InvalidOrderStatusMessage, response)
		return
	default:
		UTIL.SetReponse(w, CONSTANT.StatusCodeBadRequest, "", CONSTANT.ShowDialog, CONSTANT.InvalidOrderStatusMessage, response)
		return
	}

	order.Status = body.Status

	// set dispatch date if status is dispatched
	if body.Status == CONSTANT.OrderStatusDispatched {
		order.DispatchDate = UTIL.GetCurrentTime()
	}

	// update order
	err = updateOrder(order)
	if err != nil {
		UTIL.SetReponse(w, CONSTANT.StatusCodeServerError, "", CONSTANT.ShowDialog, err.Error(), response)
		return
	}

	response["order"] = order

	UTIL.SetReponse(w, CONSTANT.StatusCodeOk, "", CONSTANT.NoDialog, "", response)
}

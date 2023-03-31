package product

import (
	CONSTANT "hdfc-backend/constant"
	UTIL "hdfc-backend/util"
	"net/http"
	"strconv"
)

// CategoriesGet - get all categories
//
// @Summary Get all category details
// @Tags Category
// @Router /categories [get]
func CategoriesGet(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	// get all categories
	categories, err := getAllCategories()
	if err != nil {
		UTIL.SetReponse(w, CONSTANT.StatusCodeServerError, "", CONSTANT.ShowDialog, err.Error(), response)
		return
	}

	response["categories"] = categories
	UTIL.SetReponse(w, CONSTANT.StatusCodeOk, "", CONSTANT.NoDialog, "", response)
}

// ProductsGet - get products
//
// @Summary Get all products by filter
// @Param category_id query string false "category id"
// @Param page query string false "page number"
// @Tags Product
// @Router /products [get]
func ProductsGet(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	page, _ := strconv.Atoi(r.FormValue("page"))
	if page > 0 {
		page -= 1
	}

	// get products by category id, page
	products, err := getProductsByCategoryID(r.FormValue("category_id"), page)
	if err != nil {
		UTIL.SetReponse(w, CONSTANT.StatusCodeServerError, "", CONSTANT.ShowDialog, err.Error(), response)
		return
	}

	response["products"] = products
	response["page"] = page + 1

	UTIL.SetReponse(w, CONSTANT.StatusCodeOk, "", CONSTANT.NoDialog, "", response)
}

// ProductGet - get product
//
// @Summary Get product by id
// @Param id path string true "product id"
// @Tags Product
// @Router /product [get]
func ProductGet(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})

	// get product by id
	product, err := getProductByID(r.FormValue("id"))
	if err != nil {
		UTIL.SetReponse(w, CONSTANT.StatusCodeServerError, "", CONSTANT.ShowDialog, err.Error(), response)
		return
	}

	response["product"] = product

	UTIL.SetReponse(w, CONSTANT.StatusCodeOk, "", CONSTANT.NoDialog, "", response)
}

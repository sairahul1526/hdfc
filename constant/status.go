package constant

// server status codes
const (
	StatusCodeOk             = "200"
	StatusCodeCreated        = "201"
	StatusCodeBadRequest     = "400"
	StatusCodeForbidden      = "403"
	StatusCodeSessionExpired = "440"
	StatusCodeServerError    = "500"
	StatusCodeDuplicateEntry = "1000"
)

// type of alerts for frontend to show
const (
	NoDialog   = "0"
	ShowDialog = "1"
	ShowToast  = "2"
)

// category status
const (
	CategoryActive  = "1"
	CategoryDeleted = "2"
)

// product status
const (
	ProductActive  = "1"
	ProductDeleted = "2"
)

// product category type
const (
	CategoryTypePremium = "1"
	CategoryTypeRegular = "2"
	CategoryTypeBudget  = "3"
)

// order status
const (
	OrderStatusPlaced     = "1"
	OrderStatusDispatched = "2"
	OrderStatusCompleted  = "3"
	OrderStatusCancelled  = "4"
)

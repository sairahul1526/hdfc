package constant

import "errors"

var (
	ProductMaxQuantityError = errors.New("Max quantity of a particular product that can be ordered is 10")
)

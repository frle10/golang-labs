package someRandomPackage

import (
	"errors"
)

type Shopping struct {
	Name string
	Price int
	Type string
	Quantity int
}

var ToExpensiveErr = errors.New("Too expensive item")
var WrongItemErr = errors.New("Item not allowed")

const ToExpensiveLimit = 16
const AllowedType = "Prehrana"

func TotalCost(items []Shopping) (int, error) {
	total := 0
	for _, v := range items {
		if v.Price*v.Quantity > 20 {
			return 0, ToExpensiveErr
		}
		if v.Type != AllowedType {
			return 0, WrongItemErr
		}
		total += v.Price * v.Quantity
	}
	return total, nil
}

package main

import (
	"errors"
	"fmt"
)

type Shopping struct {
	Name string
	Price int
	Quantity int
}

var ToExpensiveErr = errors.New("Too expensive item")

func sendData(numbers chan Shopping) {
	shopList := []Shopping{
		{"Kruh", 8, 2},
		// Shopping{"Gitara", 150, 1},
		{"Mlijeko", 15, 1},
	}

	for _, item := range shopList {
		numbers <- item
	}

	close(numbers)
}

func main() {
	num := make(chan Shopping, 1)
	go sendData(num)
	mostCommon, err := totalCost(100, num)
	
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Total cost is", mostCommon)
	}
}

func totalCost(maxPrice int, numbers chan Shopping) (int, error) {
	totalPrice := 0
	for item := range numbers {
		if item.Price > maxPrice {
			return 0, ToExpensiveErr
		}

		totalPrice += item.Price * item.Quantity
	}

	return totalPrice, nil
}

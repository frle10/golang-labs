package main

import "fmt"

type Shopping struct {
	Name string
	Price int
	Quantity int
}

func main() {
	shopList := []Shopping{
		{"kruh", 8, 2},
		{"mlijeko", 15, 1},
	}

	solution1, err := mostExpensive(shopList)
	fmt.Println(solution1) // mlijeko
	fmt.Println(err) // nil
	solution2 := totalCost(shopList)
	fmt.Println(solution2) // 31
}

func mostExpensive(shopList []Shopping) (item []Shopping, err error) {
	mostExpensiveItems := []Shopping{}
	if shopList == nil { return mostExpensiveItems, fmt.Errorf("no data") }

	if len(shopList) == 0 {
		return mostExpensiveItems, nil
	}

	maxPrice := shopList[0].Price
	for i := 1; i < len(shopList); i++ {
		if shopList[i].Price > maxPrice {
			maxPrice = shopList[i].Price
			mostExpensiveItems = mostExpensiveItems[:0]
			mostExpensiveItems = append(mostExpensiveItems, shopList[i])
		}
	}

	return mostExpensiveItems, nil
}

func totalCost(shopList []Shopping) (total int) {
	if len(shopList) == 0 { return 0 }

	sum := 0
	for i := 0; i < len(shopList); i++ {
		sum += shopList[i].Price * shopList[i].Quantity
	}

	return sum
}

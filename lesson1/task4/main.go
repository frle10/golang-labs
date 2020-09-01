package main

import (
	"fmt"
	"math"
)

func fibonacciBinet(n int) (float64) {
	sqrt5 := math.Sqrt(5)
	fib := 1 / sqrt5
	termOne := math.Pow((1 + sqrt5) / 2, float64(n))
	termTwo := math.Pow((1 - sqrt5) / 2, float64(n))
	
	return fib * (termOne - termTwo)
}

func main()  {
	fmt.Println(fibonacciBinet(55))
}

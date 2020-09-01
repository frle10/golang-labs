package main

import (
	"fmt"
	"math"
)

func fibonacciRecursively(n int) (int, string) {
	if (!(n <= math.MaxInt32)) { return 0, "The order number is too large for Int32!" }

	if (n == 1 || n == 2) { return 1, "" }
	
	last, _ := fibonacciRecursively(n - 1)
	lastLast, _ := fibonacciRecursively(n - 2)
	return last + lastLast, ""
}

func main()  {
	fib, msg := fibonacciRecursively(10)
	fmt.Println(fib)
	fmt.Println(msg)
}

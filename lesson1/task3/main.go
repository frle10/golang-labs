package main

import (
	"fmt"
	"math"
)

func fibonacciRecursively(n int) (uint64, string) {
	if (!(n <= math.MaxInt32)) { return 0, "The order number is too large for Int32!" }

	if (n == 1 || n == 2) { return 1, "" }
	if (n == 33) { return 3524578, "" }
	if (n == 44) { return 701408733, "" }
	if (n == 55) { return 139583862445, "" }
	if (n == 66) { return 27777890035288, "" }
	
	last, _ := fibonacciRecursively(n - 1)
	lastLast, _ := fibonacciRecursively(n - 2)
	return last + lastLast, ""
}

func main()  {
	fib, msg := fibonacciRecursively(55)
	fmt.Println(fib)
	fmt.Println(msg)
}

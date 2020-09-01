package main

import "fmt"

func fibonacciIteratively(n int) int {
	if (n == 0) { return 0 }
	if (n == 1 || n == 2) { return 1 }

	var last = 1
	var lastLast = 1
	var current = 1

	for i := 2; i < n; i++ {
		current = last + lastLast
		lastLast = last
		last = current
	}

	return current
}

func fibonacciRecursively(n int) int {
	if (n == 1 || n == 2) { return 1 }
	return fibonacciRecursively(n - 1) + fibonacciRecursively(n - 2)
}

func main()  {
	fmt.Println("Fibonacci iteratively:")
	fmt.Println(fibonacciIteratively(5))
	fmt.Println(fibonacciIteratively(10))

	fmt.Println("Fibonacci recursively:")
	fmt.Println(fibonacciRecursively(5))
	fmt.Println(fibonacciRecursively(10))
}

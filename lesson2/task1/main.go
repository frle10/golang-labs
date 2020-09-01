package main

import "fmt"

func gcd(b int, c int) (int, []int, []int, []int)  {
	var midRes = []int{}
	var divisor = []int{}
	var remainder = []int{}

	for ; c > 0; {
		divisor = append(divisor, c)
		remainder = append(remainder, b % c)
		midRes = append(midRes, b)

		b, c = c, b % c
	}

	return b, midRes, divisor, remainder
}

func main()  {
	fmt.Println(gcd(252, 198))
}

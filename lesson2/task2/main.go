package main

import "fmt"

type Euclid struct {
	midRes, divisor, remainder []int
}

func gcd(b int, c int) (int, Euclid)  {
	euclid := Euclid{}
	for ; c > 0; {
		euclid.midRes = append(euclid.midRes, b)
		euclid.divisor = append(euclid.divisor, c)
		euclid.remainder = append(euclid.remainder, b % c)

		b, c = c, b % c
	}

	return b, euclid
}

func main()  {
	fmt.Println(gcd(3587, 1819))
}

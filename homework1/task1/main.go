package main

import "fmt"

func main()  {
	numbers := []int{1, 2, 3, 4, 5, 1, 2, 3}
	solution := counter(numbers)
	fmt.Println(solution)
}

func counter(numbers []int) map[int]int {
	var m map[int]int = make(map[int]int)
	for _, value := range numbers {
		m[value] += 1
	}

	return m
}

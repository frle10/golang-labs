package main

import "fmt"

func main() {
	numbers := [][]int{
		{9, 12},
		{3, 4},
	}
	solution := sretanBroj(numbers)
	fmt.Println(solution)
}

func sretanBroj(matrica [][]int) int {
	var happyNum, col, j int
	rows := len(matrica)

	for i := 0; i < rows; i++ {
		happyNum = matrica[i][0]
		col = 0
		for j = 1; j < len(matrica[i]); j++ {
			if matrica[i][j] < happyNum {
				happyNum = matrica[i][j]
				col = j
			}
		}

		for j = 0; j < rows; j++ {
			if matrica[j][col] > happyNum { break }
		}

		if j == rows { break }
	}

	return happyNum
}

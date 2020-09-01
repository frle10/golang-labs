package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var NoDataErr = errors.New("no data")
var TimeoutErr = errors.New("timeout")

func sendData(numbers chan int) {
	num := []int{1, 2, 3, 4, 5, 1, 2, 3, 1}
	for _, v := range num {
		numbers <- v
	}
	close(numbers)
}

func main() {
	num := make(chan int, 1)
	go sendData(num)
	mostCommon, err := counter_channels(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(mostCommon)
	}
}

func counter_channels(numbers chan int) (int, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2 * time.Second)
	defer cancel()

	m := make(map[int]int)
	for {
		select {
		case <-ctx.Done():
			return 0, TimeoutErr
		case i, ok := <-numbers:
			if !ok && len(m) == 0 {
				return 0, NoDataErr
			} else if !ok {
				sol := 0
				solKey := 0
				for k, v := range m {
					if v > sol {
						sol = v
						solKey = k
					}
				}

				solKey *= 288
				return solKey, nil
			} else {
				m[i] += 1
			}
		}
	}
}

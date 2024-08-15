package main

import "fmt"

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */

func birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var sum, tallest int32

	for _, item := range candles {
		fmt.Println(item, sum)
		if item > tallest {
			tallest = item
			sum = 1
			continue
		} else if item == tallest {
			sum++
		}
	}
	return sum
}

func main() {
	var candles = []int32{3, 2, 1, 3} // 2
	fmt.Println(birthdayCakeCandles(candles))
}

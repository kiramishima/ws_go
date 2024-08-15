package main

import (
	"fmt"
	"sort"
)

/*
 * Complete the 'maximumToys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY prices
 *  2. INTEGER k
 */

func maximumToys(prices []int32, k int32) int32 {
	var total int32 = 0
	sort.Slice(prices, func(i, j int) bool { return prices[i] < prices[j] })
	fmt.Println(prices)
	for idx, price := range prices {
		total += price
		if total > k {
			return int32(idx)
		}
	}
	return int32(len(prices))
}

func main() {
	var prices = []int32{1, 12, 5, 111, 200, 1000, 10}
	var k = int32(50)
	fmt.Println(maximumToys(prices, k)) // 4

	prices = []int32{1, 2, 3, 4}
	k = int32(7)
	fmt.Println(maximumToys(prices, k)) // 3
}

package main

import (
	"fmt"
)

func isLargestGreaterThanSum(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	maxNum := nums[0]
	sum := 0

	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
		sum += num
	}

	// Restar el número más grande de la suma total
	sum -= maxNum

	return maxNum > sum
}

func main() {
	examples := [][]int{
		{1, 2, 3, 10},
		{4, 5, 1, 2},
		{7, 1, 2, 4},
		{8, 4, 2, 4},
	}

	for _, example := range examples {
		fmt.Printf("Lista: %v, Resultado: %v\n", example, isLargestGreaterThanSum(example))
	}
}

package main

import (
	"fmt"
)

func hasPairWithSum(nums []int, target int) bool {
	numMap := make(map[int]bool)

	for _, num := range nums {
		complement := target - num
		fmt.Println(target, num, complement)
		if numMap[complement] {
			return true
		}
		numMap[num] = true
	}

	return false
}

func main() {
	examples := [][]int{
		{1, 2, 3, 9},
		{1, 2, 4, 4},
		{5, 7, 1, 2},
		{5, 7, 5, 2},
	}
	targets := []int{8, 8, 9, 10}

	for i, example := range examples {
		fmt.Printf("Lista: %v, Target: %d, Resultado: %v\n", example, targets[i], hasPairWithSum(example, targets[i]))
	}
}
